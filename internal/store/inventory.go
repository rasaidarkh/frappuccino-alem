package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/dto"
	"frappuccino-alem/models"
)

type InventoryRepository interface {
	CreateInventoryItem(ctx context.Context, item entity.InventoryItem) (int64, error)
	GetAllInventoryItems(ctx context.Context, pagination *dto.Pagination) ([]entity.InventoryItem, error)
	GetTotalInventoryCount(ctx context.Context) (int, error)
	GetInventoryItemById(ctx context.Context, id int64) (entity.InventoryItem, error)
	DeleteInventoryItemById(ctx context.Context, id int64) (int64, error)
	UpdateByID(ctx context.Context, id int64, updateFn func(item *entity.InventoryItem) (bool, error)) error
}

type inventoryRepository struct {
	db *sql.DB
}

func NewInventoryStore(db *sql.DB) *inventoryRepository {
	return &inventoryRepository{db}
}

func (r *inventoryRepository) CreateInventoryItem(ctx context.Context, item entity.InventoryItem) (int64, error) {
	const op = "Store.CreateInventoryItem"

	ItemModel := models.Inventory{
		ItemName:  item.ItemName,
		Quantity:  item.Quantity,
		Unit:      item.Unit,
		Price:     item.Price,
		CreatedAt: item.CreatedAt,
	}
	var id int64
	row := r.db.QueryRowContext(ctx,
		"INSERT INTO inventory (item_name,quantity,unit,price) VALUES ($1,$2,$3,$4) RETURNING id",
		ItemModel.ItemName, ItemModel.Quantity, ItemModel.Unit, ItemModel.Price)
	err := row.Scan(&id)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (r *inventoryRepository) GetAllInventoryItems(ctx context.Context, pagination *dto.Pagination) ([]entity.InventoryItem, error) {
	const op = "Store.GetAllInventoryItems"
	var items []entity.InventoryItem
	query := "SELECT * FROM inventory"

	if pagination.SortBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", pagination.SortBy)
	}

	offset := (pagination.Page - 1) * pagination.PageSize
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pagination.PageSize, offset)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.InventoryItem
		err := rows.Scan(&item.ID, &item.ItemName, &item.Quantity, &item.Unit, &item.Price, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return items, nil
}

func (r *inventoryRepository) GetTotalInventoryCount(ctx context.Context) (int, error) {
	var total int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM inventory").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *inventoryRepository) GetInventoryItemById(ctx context.Context, id int64) (entity.InventoryItem, error) {
	const op = "Store.GetInventoryItemById"
	var item entity.InventoryItem

	err := r.db.QueryRowContext(ctx, "SELECT * FROM inventory WHERE id = $1", id).Scan(
		&item.ID,
		&item.ItemName,
		&item.Quantity,
		&item.Unit,
		&item.Price,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return item, fmt.Errorf("%s: %w", op, err)
	}

	return item, nil
}

func (r *inventoryRepository) DeleteInventoryItemById(ctx context.Context, id int64) (int64, error) {
	const op = "Store.DeleteInventoryItemById"

	res, err := r.db.ExecContext(ctx, "DELETE FROM inventory WHERE id = $1", id)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return rowsAffected, nil
}

func (r *inventoryRepository) UpdateByID(ctx context.Context, id int64, updateFn func(item *entity.InventoryItem) (bool, error)) error {
	const op = "Store.UpdateInventoryItemById"
	return runInTx(r.db, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, "SELECT item_name, quantity, unit, price FROM inventory WHERE id = $1 FOR UPDATE", id)

		var itemName string
		var quantity float64
		var unit string
		var price float64
		err := row.Scan(&itemName, &quantity, &unit, &price)
		if err != nil {
			return err
		}

		item := &entity.InventoryItem{
			ID:       id,
			ItemName: itemName,
			Quantity: quantity,
			Unit:     unit,
			Price:    price,
		}

		updated, err := updateFn(item)
		if err != nil {
			return err
		}

		if !updated {
			return nil
		}

		_, err = tx.ExecContext(ctx,
			"UPDATE inventory SET item_name = $1, quantity = $2, unit = $3, price = $4, updated_at = $5 WHERE id = $6",
			item.ItemName, item.Quantity, item.Unit, item.Price, item.UpdatedAt, item.ID)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		return nil
	})
}

func runInTx(db *sql.DB, fn func(tx *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	err = fn(tx)
	if err == nil {
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit transaction: %w", err)
		}
		return nil
	}

	rollbackErr := tx.Rollback()
	if rollbackErr != nil {
		return errors.Join(err, rollbackErr)
	}

	return err
}
