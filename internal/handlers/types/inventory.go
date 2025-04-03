package types

import (
	"frappuccino-alem/internal/entity"
	"time"
)

type InventoryItemRequest struct {
	Name     *string  `json:"name"`
	Quantity *float64 `json:"quantity"`
	UnitType *string  `json:"unit"`
}

func (i InventoryItemRequest) MapToInventoryItemEntity() entity.InventoryItem {
	if i.Name == nil {
		*i.Name = "NewInventoryItemName"
	}
	if i.Quantity == nil {
		*i.Quantity = 0
	}
	if i.UnitType == nil {
		*i.UnitType = "kg"
	}
	return entity.InventoryItem{
		ID:          -1,
		Name:        *i.Name,
		Quantity:    *i.Quantity,
		Unit:        *i.UnitType,
		LastUpdated: time.Now(),
	}
}
