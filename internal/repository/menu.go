package repository

import (
	"context"
	"database/sql"
	"frappuccino-alem/models"
)

type MenuRepository struct {
	db *sql.DB
}

func NewMenuRepository(db *sql.DB) *MenuRepository {
	return &MenuRepository{db}
}

func (r *MenuRepository) CreateMenuItem(ctx context.Context, item models.MenuItem) (string, error) {
	const op = "repository.CreateMenuItem"

	// logic here ...

	return "", nil
}

func (r *MenuRepository) GetAllMenuItems(ctx context.Context) ([]models.MenuItem, error) {
	const op = "repository.GetAllMenuItems"
	var items []models.MenuItem

	// logic here ...

	return items, nil
}

func (r *MenuRepository) GetMenuItemById(ctx context.Context, id string) (models.MenuItem, error) {
	const op = "repository.GetMenuItemById"
	var item models.MenuItem

	// logic here ...

	return item, nil
}

func (r *MenuRepository) UpdateMenuItemById(ctx context.Context, id string, item models.MenuItem) error {
	const op = "repository.UpdateMenuItemById"

	// logic here ...

	return nil
}

func (r *MenuRepository) DeleteMenuItemById(ctx context.Context, id string) error {
	const op = "repository.DeleteMenuItemById"

	// logic here ...

	return nil
}
