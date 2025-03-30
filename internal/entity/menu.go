package models

import (
	"time"
)

type MenuItem struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Categories  []string
	Allergens   []string
	Metadata    JSONB
	CreatedAt   time.Time
}

type PriceHistory struct {
	ID         int
	MenuItemId int
	OldPrice   float64
	NewPrice   float64
	CreatedAt  time.Time
}
