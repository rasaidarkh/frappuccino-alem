package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/dto"
	"frappuccino-alem/models"
	"frappuccino-alem/models/mapper"

	"github.com/lib/pq"
)

type MenuRepository interface {
	CreateMenuItem(ctx context.Context, item entity.MenuItem) (int64, error)
	GetAllMenuItems(ctx context.Context, pagination *dto.Pagination) ([]entity.MenuItem, error)
	GetTotalMenuCount(ctx context.Context) (int, error)
	GetMenuItemById(ctx context.Context, id int64) (entity.MenuItem, error)
	DeleteMenuItemById(ctx context.Context, id int64) error
	UpdateByID(ctx context.Context, id int64, updateFn func(item *entity.MenuItem) (bool, error)) error
}

type menuRepository struct {
	db *sql.DB
}

func NewMenuStore(db *sql.DB) *menuRepository {
	return &menuRepository{db}
}

func (s *menuRepository) CreateMenuItem(ctx context.Context, item entity.MenuItem) (int64, error) {
	const op = "Store.CreateMenuItem"
	var id int64

	err := runInTx(s.db, func(tx *sql.Tx) error {
		modelItem := mapper.ToMenuItemModel(item)

		err := tx.QueryRowContext(ctx,
			`INSERT INTO menu_items (name, description, price, categories, allergens, metadata)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id`,
			modelItem.Name, modelItem.Description, modelItem.Price,
			modelItem.Categories, modelItem.Allergens, modelItem.Metadata,
		).Scan(&id)
		if err != nil {
			return fmt.Errorf("insert menu item: %w", err)
		}

		// Insert ingredients if present
		if len(item.Ingredients) > 0 {
			valueStrings := make([]string, 0, len(item.Ingredients))
			valueArgs := make([]interface{}, 0, len(item.Ingredients)*3)

			for i, ing := range item.Ingredients {
				valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
				valueArgs = append(valueArgs, id, ing.ItemID, ing.Quantity)
			}

			_, err = tx.ExecContext(ctx,
				fmt.Sprintf(`INSERT INTO menu_item_ingredients (menu_item_id, ingredient_id, quantity_used)  
					VALUES %s`, strings.Join(valueStrings, ",")),
				valueArgs...)
			if err != nil {
				return fmt.Errorf("insert ingredients: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *menuRepository) GetAllMenuItems(ctx context.Context, pagination *dto.Pagination) ([]entity.MenuItem, error) {
	const op = "Store.GetAllMenuItems"

	query := `
		SELECT id, name, description, price, categories, allergens, 
			metadata, created_at, updated_at 
		FROM menu_items
	`

	if pagination.SortBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", pagination.SortBy)
	}

	offset := (pagination.Page - 1) * pagination.PageSize
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pagination.PageSize, offset)

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var modelItems []models.MenuItem
	for rows.Next() {
		var model models.MenuItem
		err := rows.Scan(
			&model.ID,
			&model.Name,
			&model.Description,
			&model.Price,
			&model.Categories,
			&model.Allergens,
			&model.Metadata,
			&model.CreatedAt,
			&model.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		modelItems = append(modelItems, model)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	entities := make([]entity.MenuItem, 0, len(modelItems))
	for _, model := range modelItems {
		ingredients, err := s.getIngredientsForMenuItem(ctx, model.ID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		entities = append(entities, mapper.ToMenuItemEntity(model, ingredients))
	}

	return entities, nil
}

func (s *menuRepository) getIngredientsForMenuItem(ctx context.Context, menuItemID int64) ([]entity.MenuIngredient, error) {
	const op = "Store.getIngredientsForMenuItem"

	query := `
        SELECT 
            i.id, 
            i.item_name,
            mi.quantity_used, 
            i.unit, 
            i.price
        FROM menu_item_ingredients mi
        JOIN inventory i ON mi.ingredient_id = i.id
        WHERE mi.menu_item_id = $1
    `

	rows, err := s.db.QueryContext(ctx, query, menuItemID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var ingredients []entity.MenuIngredient
	for rows.Next() {
		var model models.Inventory
		var quantityUsed float64
		err := rows.Scan(
			&model.ID,
			&model.ItemName,
			&quantityUsed,
			&model.Unit,
			&model.Price,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		ingredients = append(ingredients, entity.MenuIngredient{
			ItemID:   model.ID,
			Name:     model.ItemName, // Map to entity field
			Quantity: quantityUsed,
			Unit:     model.Unit,
			Price:    model.Price,
		})
	}

	return ingredients, nil
}

func (s *menuRepository) GetTotalMenuCount(ctx context.Context) (int, error) {
	const op = "Store.GetTotalMenuCount"

	var total int
	err := s.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM menu_items").Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return total, nil
}

func (s *menuRepository) GetMenuItemById(ctx context.Context, id int64) (entity.MenuItem, error) {
	const op = "Store.GetMenuItemById"

	var model models.MenuItem
	err := s.db.QueryRowContext(ctx,
		`SELECT id, name, description, price, categories, allergens, 
		metadata, created_at, updated_at 
		FROM menu_items WHERE id = $1`,
		id,
	).Scan(
		&model.ID,
		&model.Name,
		&model.Description,
		&model.Price,
		&model.Categories,
		&model.Allergens,
		&model.Metadata,
		&model.CreatedAt,
		&model.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.MenuItem{}, fmt.Errorf("%s: %w", op, ErrNotFound)
		}
		return entity.MenuItem{}, fmt.Errorf("%s: %w", op, err)
	}

	ingredients, err := s.getIngredientsForMenuItem(ctx, model.ID)
	if err != nil {
		return entity.MenuItem{}, fmt.Errorf("%s: %w", op, err)
	}

	return mapper.ToMenuItemEntity(model, ingredients), nil
}

func (r *menuRepository) UpdateByID(ctx context.Context, id int64, updateFn func(item *entity.MenuItem) (bool, error)) error {
	const op = "Store.Menu.UpdateByID"
	return runInTx(r.db, func(tx *sql.Tx) error {
		// Use COALESCE to handle NULL arrays
		row := tx.QueryRowContext(ctx, `
            SELECT 
                name, 
                description, 
                price, 
                COALESCE(categories, '{}') AS categories,
                COALESCE(allergens, '{}') AS allergens,
                COALESCE(metadata, '{}') AS metadata
            FROM menu_items 
            WHERE id = $1 FOR UPDATE`, id)

		var (
			name        string
			description string
			price       float64
			categories  pq.StringArray
			allergens   pq.StringArray
			metadata    entity.JSONB
		)

		err := row.Scan(
			&name,
			&description,
			&price,
			&categories,
			&allergens,
			&metadata,
		)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		item := &entity.MenuItem{
			ID:          id,
			Name:        name,
			Description: description,
			Price:       price,
			Categories:  categories, // Convert to regular slice
			Allergens:   allergens,
			Metadata:    metadata,
		}

		updated, err := updateFn(item)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		if !updated {
			return nil
		}

		// Add your update execution logic here
		_, err = tx.ExecContext(ctx, `
            UPDATE menu_items SET
                name = $1,
                description = $2,
                price = $3,
                categories = $4,
                allergens = $5,
                metadata = $6,
                updated_at = NOW()
            WHERE id = $7`,
			item.Name,
			item.Description,
			item.Price,
			pq.Array(item.Categories), // Use pq.Array for updates
			pq.Array(item.Allergens),
			item.Metadata,
			id,
		)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		return nil
	})
}

// func (s *menuRepository) UpdateByID(ctx context.Context, id int64, item entity.MenuItem) error {
// 	const op = "Store.UpdateMenuItemById"

// 	tx, err := s.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}
// 	defer func() {
// 		if err != nil {
// 			tx.Rollback()
// 		}
// 	}()

// 	// First check if menu item exists
// 	var exists bool
// 	err = tx.QueryRowContext(ctx,
// 		"SELECT EXISTS(SELECT 1 FROM menu_items WHERE id = $1)",
// 		id,
// 	).Scan(&exists)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}
// 	if !exists {
// 		return fmt.Errorf("%s: %w", op, ErrNotFound)
// 	}

// 	modelItem := mapper.ToMenuItemModel(item)
// 	result, err := tx.ExecContext(ctx, `
//         UPDATE menu_items
//         SET name = $1, description = $2, price = $3,
//             categories = $4, allergens = $5, metadata = $6,
//             updated_at = NOW()
//         WHERE id = $7`,
// 		modelItem.Name,
// 		modelItem.Description,
// 		modelItem.Price,
// 		modelItem.Categories,
// 		modelItem.Allergens,
// 		modelItem.Metadata,
// 		id,
// 	)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	// Verify that the update affected at least one row
// 	rowsAffected, err := result.RowsAffected()
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}
// 	if rowsAffected == 0 {
// 		return fmt.Errorf("%s: %w", op, ErrNotFound)
// 	}

// 	// Rest of the transaction remains the same...
// 	_, err = tx.ExecContext(ctx, `
//         DELETE FROM menu_item_ingredients
//         WHERE menu_item_id = $1`,
// 		id,
// 	)
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	if len(item.Ingredients) > 0 {
// 		query := `INSERT INTO menu_item_ingredients (menu_item_id, ingredient_id, quantity_used) VALUES `
// 		values := make([]string, 0, len(item.Ingredients))
// 		args := make([]interface{}, 0, len(item.Ingredients)*3)

// 		for i, ing := range item.Ingredients {
// 			if ing.Quantity <= 0 {
// 				return fmt.Errorf("%s: quantity_used must be greater than 0", op)
// 			}
// 			values = append(values, fmt.Sprintf("($%d, $%d, $%d)", i*3+1, i*3+2, i*3+3))
// 			args = append(args, id, ing.ID, ing.Quantity)
// 		}

// 		_, err = tx.ExecContext(ctx, query+strings.Join(values, ","), args...)
// 		if err != nil {
// 			return fmt.Errorf("%s: %w", op, err)
// 		}
// 	}

// 	if err = tx.Commit(); err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	return nil
// }

func (s *menuRepository) DeleteMenuItemById(ctx context.Context, id int64) error {
	const op = "Store.DeleteMenuItemById"

	result, err := s.db.ExecContext(ctx, "DELETE FROM menu_items WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, ErrNotFound)
	}

	return nil
}
