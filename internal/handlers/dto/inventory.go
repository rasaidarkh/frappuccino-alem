package dto

import (
	"frappuccino-alem/internal/entity"
)

type InventoryItemRequest struct {
	Name     *string  `json:"name"`
	Quantity *float64 `json:"quantity"`
	UnitType *string  `json:"unit"`
	Price    *float64 `json:"price"`
}

func (r InventoryItemRequest) MapToInventoryItemEntity() entity.InventoryItem {
	return entity.InventoryItem{
		ItemName: *r.Name,
		Quantity: *r.Quantity,
		Unit:     *r.UnitType,
		Price:    *r.Price,
	}
}

type LefOverItem struct {
	Name     string  `json:"name"`
	Quantity float64 `json:"quantity"`
	UnitType string  `json:"unit"`
	Price    float64 `json:"price"`
}
