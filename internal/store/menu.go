package store

import (
	"context"
	"database/sql"
	"fmt"
	"frappuccino-alem/internal/handlers/types"
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

func (r *MenuStore) GetAllMenuItems(ctx context.Context, pagination *types.Pagination) ([]models.MenuItem, error) {
	var items []models.MenuItem
	query := "SELECT * FROM menu_items"

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
		var item models.MenuItem
		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Categories, &item.Allergens, &item.Metadata, &item.CreatedAt, &item.UpdatedAt)
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

func (r *MenuStore) GetTotalMenuCount(ctx context.Context) (int, error) {
	var total int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM menu_items").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
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
