package models

import (
	"time"
)

type Inventory struct {
	ID          int
	Name        string
	Quantity    float64
	Unit        string
	LastUpdated time.Time
}

type InventoryTransaction struct {
	ID              int
	InventoryID     int
	QuantityChange  float64
	Reason          string
	TransactionTime time.Time
}

type MenuItemIngredient struct {
	MenuItemID   int
	IngredientID int
	QuantityUsed float64
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
