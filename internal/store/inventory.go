package store

import (
	"context"
	"database/sql"
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
		UnitType:    item.Unit,
		LastUpdated: item.LastUpdated,
	}
	var id int64
	row := r.db.QueryRowContext(ctx, "INSERT INTO inventory (item_name,quantity,unit,last_updated) VALUES ($1,$2,$3,$4) RETURNING id", ItemModel.Name, ItemModel.Quantity, ItemModel.UnitType, ItemModel.LastUpdated)
	err := row.Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *InventoryStore) GetAllInventoryItems(ctx context.Context) ([]entity.InventoryItem, error) {
	const op = "Store.GetAllInventoryItems"
	var items []entity.InventoryItem
	// logic here ...

	return items, nil
}

func (r *InventoryStore) GetInventoryItemById(ctx context.Context, id int64) (entity.InventoryItem, error) {
	const op = "Store.GetInventoryItemById"
	var item entity.InventoryItem

	// logic here ...

	return item, nil
}

func (r *InventoryStore) DeleteInventoryItemById(ctx context.Context, id int64) error {
	const op = "Store.DeleteInventoryItemById"

	// logic here ...

	return nil
}

func (r *InventoryStore) UpdateInventoryItemById(ctx context.Context, id int64, item entity.InventoryItem) error {
	const op = "Store.UpdateInventoryItemById"

	// logic here ...

	return nil
}
