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
	const op = "Store.GetAllInventoryItems"
	var items []entity.InventoryItem

	var totalItems int
	countQuery := "SELECT COUNT(*) FROM inventory"
	err := r.db.QueryRowContext(ctx, countQuery).Scan(&totalItems)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to count total items: %w", op, err)
	}

	maxPages := (totalItems + pagination.PageSize - 1) / pagination.PageSize
	if pagination.Page > maxPages {
		return nil, fmt.Errorf("%s: requested page %d exceeds maximum page number %d", op, pagination.Page, maxPages)
	}

	query := "SELECT * FROM inventory"
	if pagination.SortBy != "" {
		switch pagination.SortBy {
		case types.SortByID:
			query += " ORDER BY id"
		case types.SortByQuantity:
			query += " ORDER BY quantity"
		case types.SortByName:
			query += " ORDER BY name"
		case types.SortByDate:
			query += " ORDER BY last_updated"
		default:
			return nil, fmt.Errorf("invalid sort option: %s", pagination.SortBy)
		}
	}

	if pagination.Page > 0 && pagination.PageSize > 0 {
		offset := (pagination.Page - 1) * pagination.PageSize
		query += fmt.Sprintf(" LIMIT %d OFFSET %d", pagination.PageSize, offset)
	}

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return items, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return items, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var item entity.InventoryItem
		err := rows.Scan(&item.ID, &item.Name, &item.Quantity, &item.Unit, &item.LastUpdated)
		if err != nil {
			return items, fmt.Errorf("%s: %w", op, err)
		}
		items = append(items, item)
	}

	err = rows.Err()
	if err != nil {
		return items, fmt.Errorf("%s: %w", op, err)
	}

	return items, nil
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
