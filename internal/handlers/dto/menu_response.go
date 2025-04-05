package dto

import (
	"frappuccino-alem/internal/entity"
	"time"
)

type MenuItemResponse struct {
	ID          int64                        `json:"id"`
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Price       float64                      `json:"price"`
	Categories  []string                     `json:"categories"`
	Allergens   []string                     `json:"allergens"`
	Metadata    map[string]any               `json:"metadata"`
	Ingredients []MenuItemIngredientResponse `json:"ingredients"`
	CreatedAt   string                       `json:"created_at"`
	UpdatedAt   string                       `json:"updated_at"`
}

type MenuItemIngredientResponse struct {
	InventoryID  int64   `json:"inventory_id"`
	Name         string  `json:"name"`
	QuantityUsed float64 `json:"quantity_used"`
}

type CreatedResponse struct {
	ID int64 `json:"id"`
}

func ToMenuItemResponse(e entity.MenuItem) MenuItemResponse {
	return MenuItemResponse{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Price:       e.Price,
		Categories:  e.Categories,
		Allergens:   e.Allergens,
		Metadata:    e.Metadata,
		Ingredients: toIngredientResponses(e.Ingredients),
		CreatedAt:   formatTime(e.CreatedAt),
		UpdatedAt:   formatTime(e.UpdatedAt),
	}
}

func toIngredientResponses(ingredients []entity.MenuItemIngredient) []MenuItemIngredientResponse {
	res := make([]MenuItemIngredientResponse, len(ingredients))
	for i, ing := range ingredients {
		res[i] = MenuItemIngredientResponse{
			InventoryID:  ing.InventoryID,
			Name:         ing.Name,
			QuantityUsed: ing.QuantityUsed,
		}
	}
	return res
}

func formatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
