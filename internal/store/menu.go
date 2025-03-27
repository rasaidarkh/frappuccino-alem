package store

import (
	"context"
	"database/sql"
	"frappuccino-alem/models"
)

type MenuStore struct {
	db *sql.DB
}

func NewMenuStore(db *sql.DB) *MenuStore {
	return &MenuStore{db}
}

func (r *MenuStore) CreateMenuItem(ctx context.Context, item models.MenuItem) (string, error) {
	const op = "Store.CreateMenuItem"

	// logic here ...

	return "", nil
}

func (r *MenuStore) GetAllMenuItems(ctx context.Context) ([]models.MenuItem, error) {
	const op = "Store.GetAllMenuItems"
	var items []models.MenuItem

	// logic here ...

	return items, nil
}

func (r *MenuStore) GetMenuItemById(ctx context.Context, id string) (models.MenuItem, error) {
	const op = "Store.GetMenuItemById"
	var item models.MenuItem

	// logic here ...

	return item, nil
}

func (r *MenuStore) UpdateMenuItemById(ctx context.Context, id string, item models.MenuItem) error {
	const op = "Store.UpdateMenuItemById"

	// logic here ...

	return nil
}

func (r *MenuStore) DeleteMenuItemById(ctx context.Context, id string) error {
	const op = "Store.DeleteMenuItemById"

	// logic here ...

	return nil
}
