package store

import (
	"context"
	"database/sql"
	"frappuccino-alem/models"
)

type InventoryStore struct {
	db *sql.DB
}

func NewInventoryStore(db *sql.DB) *InventoryStore {
	return &InventoryStore{db}
}

func (r *InventoryStore) CreateInventoryItem(ctx context.Context, item models.InventoryItem) (string, error) {
	const op = "Store.CreateInventoryItem"

	// logic here ...

	return "", nil
}

func (r *InventoryStore) GetAllInventoryItems(ctx context.Context) ([]models.InventoryItem, error) {
	const op = "Store.GetAllInventoryItems"
	var items []models.InventoryItem
	// logic here ...

	return items, nil
}

func (r *InventoryStore) GetInventoryItemById(ctx context.Context, id string) (models.InventoryItem, error) {
	const op = "Store.GetInventoryItemById"
	var item models.InventoryItem

	// logic here ...

	return item, nil
}

func (r *InventoryStore) DeleteInventoryItemById(ctx context.Context, id string) error {
	const op = "Store.DeleteInventoryItemById"

	// logic here ...

	return nil
}

func (r *InventoryStore) UpdateInventoryItemById(ctx context.Context, id string, item models.InventoryItem) error {
	const op = "Store.UpdateInventoryItemById"

	// logic here ...

	return nil
}
