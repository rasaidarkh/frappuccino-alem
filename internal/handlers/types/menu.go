package types

import (
	"frappuccino-alem/internal/entity"
	"time"
)

type MenuItemCreateRequest struct {
	Name        string                  `json:"name" validate:"required"`
	Description string                  `json:"description"`
	Price       float64                 `json:"price" validate:"required,gt=0"`
	Categories  []string                `json:"categories"`
	Allergens   []string                `json:"allergens"`
	Metadata    map[string]interface{}  `json:"metadata"`
	Ingredients []MenuItemIngredientDTO `json:"ingredients" validate:"required,gt=0"`
}

type MenuItemUpdateRequest struct {
	Name        *string                  `json:"name"`
	Description *string                  `json:"description"`
	Price       *float64                 `json:"price"`
	Categories  *[]string                `json:"categories"`
	Allergens   *[]string                `json:"allergens"`
	Metadata    *map[string]interface{}  `json:"metadata"`
	Ingredients *[]MenuItemIngredientDTO `json:"ingredients"`
}

type MenuItemIngredientDTO struct {
	InventoryID  int64   `json:"inventory_id" validate:"required"`
	QuantityUsed float64 `json:"quantity_used" validate:"required,gt=0"`
}

// Response DTOs
type MenuItemResponse struct {
	ID          int64                        `json:"id"`
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Price       float64                      `json:"price"`
	Categories  []string                     `json:"categories"`
	Allergens   []string                     `json:"allergens"`
	Metadata    map[string]interface{}       `json:"metadata"`
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

// Conversion methods
func (dto *MenuItemCreateRequest) ToEntity() entity.MenuItem {
	return entity.MenuItem{
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Categories:  dto.Categories,
		Allergens:   dto.Allergens,
		Metadata:    dto.Metadata,
		Ingredients: mapIngredientsDTOToEntity(dto.Ingredients),
	}
}

func (dto *MenuItemUpdateRequest) ToEntity() entity.MenuItem {
	item := entity.MenuItem{}

	if dto.Name != nil {
		item.Name = *dto.Name
	}
	if dto.Description != nil {
		item.Description = *dto.Description
	}
	if dto.Price != nil {
		item.Price = *dto.Price
	}
	if dto.Categories != nil {
		item.Categories = *dto.Categories
	}
	if dto.Allergens != nil {
		item.Allergens = *dto.Allergens
	}
	if dto.Metadata != nil {
		item.Metadata = *dto.Metadata
	}
	if dto.Ingredients != nil {
		item.Ingredients = mapIngredientsDTOToEntity(*dto.Ingredients)
	}

	return item
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

// Helper functions
func mapIngredientsDTOToEntity(dtos []MenuItemIngredientDTO) []entity.InventoryItem {
	ingredients := make([]entity.InventoryItem, len(dtos))
	for i, dto := range dtos {
		ingredients[i] = entity.InventoryItem{
			ID:           dto.InventoryID,
			QuantityUsed: dto.QuantityUsed,
		}
	}
	return ingredients
}

func mapIngredientsEntityToResponse(entities []entity.InventoryItem) []MenuItemIngredientResponse {
	response := make([]MenuItemIngredientResponse, len(entities))
	for i, e := range entities {
		response[i] = MenuItemIngredientResponse{
			InventoryID:  e.ID,
			Name:         e.ItemName,
			QuantityUsed: e.QuantityUsed,
			Unit:         e.Unit,
		}
	}
	return response
}
