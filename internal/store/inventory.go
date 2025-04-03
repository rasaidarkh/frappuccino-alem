package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/types"
	"frappuccino-alem/models"
)

type InventoryStore struct {
	db *sql.DB
}

func NewInventoryStore(db *sql.DB) *InventoryStore {
	return &InventoryStore{db}
}

func (r *InventoryStore) CreateInventoryItem(ctx context.Context, item entity.InventoryItem) (int64, error) {
	const op = "Store.CreateInventoryItem"

	ItemModel := models.Inventory{
		Name:        item.Name,
		Quantity:    item.Quantity,
		Unit:        item.Unit,
		LastUpdated: item.LastUpdated,
	}
	var id int64
	row := r.db.QueryRowContext(ctx, "INSERT INTO inventory (item_name,quantity,unit,last_updated) VALUES ($1,$2,$3,$4) RETURNING id", ItemModel.Name, ItemModel.Quantity, ItemModel.Unit, ItemModel.LastUpdated)
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *InventoryStore) GetAllInventoryItems(ctx context.Context, pagination *types.Pagination) ([]entity.InventoryItem, error) {
	var items []entity.InventoryItem
	query := "SELECT * FROM inventory"

	if pagination.SortBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", pagination.SortBy)
	}

	offset := (pagination.Page - 1) * pagination.PageSize
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pagination.PageSize, offset)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.InventoryItem
		err := rows.Scan(&item.ID, &item.Name, &item.Quantity, &item.Unit, &item.LastUpdated)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *InventoryStore) GetTotalInventoryCount(ctx context.Context) (int, error) {
	var total int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM inventory").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *InventoryStore) GetInventoryItemById(ctx context.Context, id int64) (entity.InventoryItem, error) {
	const op = "Store.GetInventoryItemById"
	var item entity.InventoryItem

	err := r.db.QueryRowContext(ctx, "SELECT * FROM inventory WHERE id = $1", id).Scan(&item.ID, &item.Name, &item.Quantity, &item.Unit, &item.LastUpdated)
	if err != nil {
		return item, fmt.Errorf("%s: %w", op, err)
	}

	return item, nil
}

func (r *InventoryStore) DeleteInventoryItemById(ctx context.Context, id int64) (int64, error) {
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

func (r *InventoryStore) UpdateInventoryItemById(ctx context.Context, id int64, item entity.InventoryItem) (int64, error) {
	const op = "Store.UpdateInventoryItemById"

	res, err := r.db.ExecContext(ctx,
		"UPDATE inventory SET item_name = $1, quantity = $2, unit = $3, last_updated = $4 WHERE id = $5",
		item.Name, item.Quantity, item.Unit, item.LastUpdated, id)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return rowsAffected, nil
}

func (r *InventoryStore) UpdateByID(ctx context.Context, id int64, updateFn func(item *entity.InventoryItem) (bool, error)) error {
	const op = "Store.UpdateInventoryItemById"
	return runInTx(r.db, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, "SELECT item_name, quantity, unit FROM inventory WHERE id = $1 FOR UPDATE", id)

		var itemName string
		var quantity float64
		var unit string
		err := row.Scan(&itemName, &quantity, &unit)
		if err != nil {
			return err
		}

		item := &entity.InventoryItem{
			ID:       int(id),
			Name:     itemName,
			Quantity: quantity,
			Unit:     unit,
		}

		updated, err := updateFn(item)
		if err != nil {
			return err
		}

		if !updated {
			return nil
		}

		_, err = tx.ExecContext(ctx, "UPDATE inventory SET item_name = $1, quantity = $2, unit = $3, last_updated = $4 WHERE id = $5",
			item.Name, item.Quantity, item.Unit, item.LastUpdated, item.ID)
		if err != nil {
			return err
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
