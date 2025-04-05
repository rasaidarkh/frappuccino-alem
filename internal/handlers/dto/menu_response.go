package dto

import (
	"time"

	"frappuccino-alem/internal/entity"
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
	Unit         string  `json:"unit"`
}

type CreatedResponse struct {
	ID int64 `json:"id"`
}

func FromEntity(entity entity.MenuItem) MenuItemResponse {
	return MenuItemResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Price:       entity.Price,
		Categories:  entity.Categories,
		Allergens:   entity.Allergens,
		Metadata:    entity.Metadata,
		Ingredients: mapIngredientsEntityToResponse(entity.Ingredients),
		CreatedAt:   entity.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   entity.UpdatedAt.Format(time.RFC3339),
	}
}

func mapIngredientsEntityToResponse(entities []entity.InventoryItem) []MenuItemIngredientResponse {
	response := make([]MenuItemIngredientResponse, len(entities))
	for i, e := range entities {
		response[i] = MenuItemIngredientResponse{
			InventoryID:  e.ID,
			Name:         e.ItemName,
			QuantityUsed: e.Quantity,
			Unit:         e.Unit,
		}
	}
	return response
}

func ToMenuItemResponse(entity entity.MenuItem) MenuItemResponse {
	return MenuItemResponse{
		ID:          entity.ID,
		Name:        entity.Name,
		Description: entity.Description,
		Price:       entity.Price,
		Categories:  entity.Categories,
		Allergens:   entity.Allergens,
		Metadata:    entity.Metadata,
		Ingredients: ToMenuItemIngredientResponses(entity.Ingredients),
		CreatedAt:   formatTime(entity.CreatedAt),
		UpdatedAt:   formatTime(entity.UpdatedAt),
	}
}

func ToMenuItemIngredientResponses(ingredients []entity.InventoryItem) []MenuItemIngredientResponse {
	responses := make([]MenuItemIngredientResponse, len(ingredients))
	for i, ing := range ingredients {
		responses[i] = ToMenuItemIngredientResponse(ing)
	}
	return responses
}

func ToMenuItemIngredientResponse(ingredient entity.InventoryItem) MenuItemIngredientResponse {
	return MenuItemIngredientResponse{
		InventoryID:  ingredient.ID,
		Name:         ingredient.ItemName,
		QuantityUsed: ingredient.Quantity,
		Unit:         ingredient.Unit,
	}
}

func formatTime(t time.Time) string {
	return t.Format(time.RFC3339)
}
