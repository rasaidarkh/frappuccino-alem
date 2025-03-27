package models

import (
	"database/sql"
	"time"
)

type Inventory struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	StockLevel  float64      `json:"stock_level"` // Supports fractional amounts (e.g., 0.5 kg)
	UnitType    string       `json:"unit_type"`   // "kg", "liters", "pieces", etc.
	LastUpdated sql.NullTime `json:"last_updated,omitempty"`
	CreatedAt   time.Time    `json:"created_at"`
}

type MenuItemIngredient struct {
	ID               int64   `json:"id"`
	MenuItemID       int64   `json:"menu_item_id"`
	InventoryID      int64   `json:"inventory_id"`
	QuantityRequired float64 `json:"quantity_required"` // Amount needed per recipe
}

type InventoryTransaction struct {
	ID              int64      `json:"id"`
	InventoryID     int64      `json:"inventory_id"`
	ChangeType      ChangeType `json:"change_type"`      // ENUM: "restock", "usage", "waste"
	QuantityChanged float64    `json:"quantity_changed"` // Amount added or removed
	ChangedAt       time.Time  `json:"changed_at"`
}

type ChangeType int

const (
	TypeRestock ChangeType = iota
	TypeUsage
	TypeWaste
)

func (c ChangeType) String() string {
	switch c {
	case TypeRestock:
		return "restock"
	case TypeUsage:
		return "usage"
	case TypeWaste:
		return "waste"
	default:
		return "unknown"
	}
}

func (c ChangeType) IsValid() bool {
	switch c {
	case TypeRestock, TypeUsage, TypeWaste:
		return true
	}
	return false
}
