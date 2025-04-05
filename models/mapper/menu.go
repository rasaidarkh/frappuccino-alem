// mapper/menu.go
package mapper

import (
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/models"

	"github.com/lib/pq"
)

func ToMenuItemModel(e entity.MenuItem) models.MenuItem {
	return models.MenuItem{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		Categories:  pq.StringArray(e.Categories),
		Allergens:   pq.StringArray(e.Allergens),
		Metadata:    models.JSONB(e.Metadata),
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

func ToMenuItemEntity(m models.MenuItem, ingredients []entity.InventoryItem) entity.MenuItem {
	return entity.MenuItem{
		ID:          m.ID,
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		Categories:  []string(m.Categories),
		Allergens:   []string(m.Allergens),
		Metadata:    entity.JSONB(m.Metadata),
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
		Ingredients: ingredients,
	}
}

func ToMenuItemIngredientModel(e entity.InventoryItem, menuItemID int) models.MenuItemIngredient {
	return models.MenuItemIngredient{
		MenuItemID:   int64(menuItemID),
		InventoryID:  int64(e.ID),
		QuantityUsed: e.Quantity,
	}
}

func ToInventoryItemEntity(m models.Inventory) entity.InventoryItem {
	return entity.InventoryItem{
		ID:        m.ID,
		ItemName:  m.ItemName,
		Quantity:  0,
		Unit:      m.Unit,
		Price:     m.Price,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
