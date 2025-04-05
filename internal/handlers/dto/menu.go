package dto

import (
	"fmt"
	"strconv"
	"time"

	"frappuccino-alem/internal/entity"
)

type MenuItemRequest struct {
	Name        *string                 `json:"name"`
	Description *string                 `json:"description"`
	Price       *float64                `json:"price"`
	Categories  *[]string               `json:"categories"`
	Allergens   *[]string               `json:"allergens"`
	Metadata    *map[string]interface{} `json:"metadata"`
	Ingredients *[]IngredientRequest    `json:"ingredients"`
}

type IngredientRequest struct {
	ItemID   *int64   `json:"item_id"`
	Quantity *float64 `json:"quantity"`
}

func (r MenuItemRequest) Validate() error {
	if r.Ingredients == nil {
		return fmt.Errorf("at least one ingredient is required")
	}

	for _, ing := range *r.Ingredients {
		if ing.ItemID == nil {
			return fmt.Errorf("invalid ingredient property: item_id is required")
		}
		if ing.Quantity == nil {
			return fmt.Errorf("invalid ingredient property: quantity is required")
		}
		if *ing.Quantity <= 0 {
			return fmt.Errorf("invalid quantity: must be greater than 0")
		}
	}

	return nil
}

func (r MenuItemRequest) MapToEntity() entity.MenuItem {
	ingredients := make([]entity.InventoryItem, 0)
	if r.Ingredients != nil {
		for _, i := range *r.Ingredients {
			ingredients = append(ingredients, entity.InventoryItem{
				ID:       *i.ItemID,
				Quantity: *i.Quantity,
			})
		}
	}

	return entity.MenuItem{
		Name:        *r.Name,
		Description: *r.Description,
		Price:       *r.Price,
		Categories:  *r.Categories,
		Allergens:   *r.Allergens,
		Metadata:    *r.Metadata,
		Ingredients: ingredients,
	}
}

type MenuIngredientResponse struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	Unit     string  `json:"unit"`
	Price    float64 `json:"price"`
}

type MenuItemResponse struct {
	ID          string                   `json:"id"`
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Price       float64                  `json:"price"`
	Categories  []string                 `json:"categories"`
	Allergens   []string                 `json:"allergens"`
	Metadata    map[string]interface{}   `json:"metadata"`
	Ingredients []MenuIngredientResponse `json:"ingredients"`
}

type MenuItemDetailedResponse struct {
	MenuItemResponse
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func MenuItemToResponse(m entity.MenuItem) MenuItemResponse {
	ingredients := make([]MenuIngredientResponse, 0)
	for _, i := range m.Ingredients {
		ingredients = append(ingredients, MenuIngredientResponse{
			ID:       strconv.FormatInt(i.ID, 10),
			Name:     i.ItemName,
			Quantity: i.Quantity,
			Unit:     i.Unit,
			Price:    i.Price,
		})
	}

	return MenuItemResponse{
		ID:          strconv.FormatInt(m.ID, 10),
		Name:        m.Name,
		Description: m.Description,
		Price:       m.Price,
		Categories:  m.Categories,
		Allergens:   m.Allergens,
		Metadata:    m.Metadata,
		Ingredients: ingredients,
	}
}

func MenuItemToDetailedResponse(m entity.MenuItem) MenuItemDetailedResponse {
	ingredients := make([]MenuIngredientResponse, 0)
	for _, i := range m.Ingredients {
		ingredients = append(ingredients, MenuIngredientResponse{
			ID:       strconv.FormatInt(i.ID, 10),
			Name:     i.ItemName,
			Quantity: i.Quantity,
			Unit:     i.Unit,
			Price:    i.Price,
		})
	}

	return MenuItemDetailedResponse{
		MenuItemResponse: MenuItemToResponse(m),
		CreatedAt:        m.CreatedAt.Format(time.RFC3339),
		UpdatedAt:        m.UpdatedAt.Format(time.RFC3339),
	}
}

func MenuItemsToResponse(items []entity.MenuItem) []MenuItemResponse {
	res := make([]MenuItemResponse, 0, len(items))
	for _, m := range items {
		res = append(res, MenuItemToResponse(m))
	}
	return res
}
