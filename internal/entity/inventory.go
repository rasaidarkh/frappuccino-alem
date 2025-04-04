package entity

import (
	"time"
)

type InventoryItem struct {
	ID           int64
	ItemName     string
	QuantityUsed float64
	Unit         string
	Price        float64
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
