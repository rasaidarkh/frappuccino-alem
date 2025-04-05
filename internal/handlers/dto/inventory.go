package dto

import (
	"strconv"
	"time"

	"frappuccino-alem/internal/entity"
)

type InventoryItemRequest struct {
	Name     *string  `json:"name"`
	Quantity *float64 `json:"quantity"`
	UnitType *string  `json:"unit"`
	Price    *float64 `json:"price"`
}

func (r InventoryItemRequest) MapToEntity() entity.InventoryItem {
	return entity.InventoryItem{
		ItemName: *r.Name,
		Quantity: *r.Quantity,
		Unit:     *r.UnitType,
		Price:    *r.Price,
	}
}

type InventoryItemResponse struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Quantity  float64 `json:"quantity"`
	Unit      string  `json:"unit"`
	Price     float64 `json:"price"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func InventoryItemToResponse(e entity.InventoryItem) InventoryItemResponse {
	return InventoryItemResponse{
		ID:        strconv.FormatInt(e.ID, 10),
		Name:      e.ItemName,
		Quantity:  e.Quantity,
		Unit:      e.Unit,
		Price:     e.Price,
		CreatedAt: e.CreatedAt.Format(time.RFC3339),
		UpdatedAt: e.UpdatedAt.Format(time.RFC3339),
	}
}

type LeftOverItem struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	UnitType string  `json:"unit"`
	Price    float64 `json:"price"`
}

func InventoryItemToLeftOver(e entity.InventoryItem) LeftOverItem {
	return LeftOverItem{
		Name:     e.ItemName,
		Quantity: e.Quantity,
		UnitType: e.Unit,
		Price:    e.Price,
	}
}
