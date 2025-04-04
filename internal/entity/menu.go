package entity

import (
	"time"
)

type MenuItem struct {
	ID          int64
	Name        string
	Description string
	Price       float64
	Categories  []string
	Allergens   []string
	Metadata    JSONB
	Ingredients []InventoryItem
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// type PriceHistory struct {
// 	ID         int
// 	MenuItemId int
// 	OldPrice   float64
// 	NewPrice   float64
// 	CreatedAt  time.Time
// }
