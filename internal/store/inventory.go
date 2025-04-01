package store

import (
	"context"
	"database/sql"
	"fmt"
	"frappuccino-alem/internal/entity"
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

func (r *InventoryStore) GetAllInventoryItems(ctx context.Context) ([]entity.InventoryItem, error) {
	const op = "Store.GetAllInventoryItems"
	var items []entity.InventoryItem

	stmt, err := r.db.PrepareContext(ctx, "SELECT * FROM inventory")
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

	res, err := r.db.ExecContext(ctx, "DELETE FROM inventroy WHERE id = $1", id)
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
