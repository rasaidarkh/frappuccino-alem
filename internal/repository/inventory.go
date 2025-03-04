package repository

import (
	"context"
	"database/sql"
	"frappuccino-alem/models"
)

type InventoryRepository struct {
	db *sql.DB
}

func NewInventoryRepository(db *sql.DB) *InventoryRepository {
	return &InventoryRepository{db}
}

func (r *InventoryRepository) CreateInventoryItem(ctx context.Context, item models.InventoryItem) (string, error) {
	const op = "repository.CreateInventoryItem"

	// logic here ...

	return "", nil
}

func (r *InventoryRepository) GetAllInventoryItems(ctx context.Context) ([]models.InventoryItem, error) {
	const op = "repository.GetAllInventoryItems"
	var items []models.InventoryItem

	// logic here ...

	return items, nil
}

func (r *InventoryRepository) GetInventoryItemById(ctx context.Context, id string) (models.InventoryItem, error) {
	const op = "repository.GetInventoryItemById"
	var item models.InventoryItem

	// logic here ...

	return item, nil
}

func (r *InventoryRepository) DeleteInventoryItemById(ctx context.Context, id string) error {
	const op = "repository.DeleteInventoryItemById"

	// logic here ...

	return nil
}

func (r *InventoryRepository) UpdateInventoryItemById(ctx context.Context, id string, item models.InventoryItem) error {
	const op = "repository.UpdateInventoryItemById"

	// logic here ...

	return nil
}
