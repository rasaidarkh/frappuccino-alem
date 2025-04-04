package models

import (
	"time"
)

type Inventory struct {
	ID        int64   `json:"id"`
	ItemName  string  `json:"item_name"`
	Quantity  float64 `json:"quantity"` // Supports fractional amounts (e.g., 0.5 kg)
	Unit      string  `json:"unit"`     // "kg", "liters", "pieces", etc.
	Price     float64
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type InventoryTransaction struct {
	ID              int64     `json:"id"`
	InventoryID     int64     `json:"inventory_id"`
	ChangeType      string    `json:"change_type"`      // ENUM: "restock", "usage", "waste"
	QuantityChanged float64   `json:"quantity_changed"` // Amount added or removed
	ChangedAt       time.Time `json:"changed_at"`
}
