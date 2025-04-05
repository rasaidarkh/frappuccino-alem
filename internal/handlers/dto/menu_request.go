package dto

import (
	"frappuccino-alem/internal/entity"
)

type MenuItemCreateRequest struct {
	Name        string                  `json:"name" `
	Description string                  `json:"description"`
	Price       float64                 `json:"price"`
	Categories  []string                `json:"categories"`
	Allergens   []string                `json:"allergens"`
	Metadata    map[string]any          `json:"metadata"`
	Ingredients []MenuItemIngredientDTO `json:"ingredients"`
}

type MenuItemUpdateRequest struct {
	Name        *string                  `json:"name"`
	Description *string                  `json:"description"`
	Price       *float64                 `json:"price"`
	Categories  *[]string                `json:"categories"`
	Allergens   *[]string                `json:"allergens"`
	Metadata    *map[string]any          `json:"metadata"`
	Ingredients *[]MenuItemIngredientDTO `json:"ingredients"`
}

type MenuItemIngredientDTO struct {
	InventoryID  int64   `json:"inventory_id" `
	QuantityUsed float64 `json:"quantity_used"`
}

// Response DTOs

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

func (r *MenuItemUpdateRequest) IsEmpty() bool {
	return r.Name == nil &&
		r.Description == nil &&
		r.Price == nil &&
		r.Categories == nil &&
		r.Allergens == nil &&
		r.Metadata == nil &&
		r.Ingredients == nil
}

// Helper functions
func mapIngredientsDTOToEntity(dtos []MenuItemIngredientDTO) []entity.InventoryItem {
	ingredients := make([]entity.InventoryItem, len(dtos))
	for i, dto := range dtos {
		ingredients[i] = entity.InventoryItem{
			ID:       dto.InventoryID,
			Quantity: dto.QuantityUsed,
		}
	}
	return ingredients
}
