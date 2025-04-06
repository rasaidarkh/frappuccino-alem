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
	Ingredients []MenuIngredient
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type MenuIngredient struct {
	ItemID   int64
	Name     string
	Quantity float64
	Unit     string
	Price    float64
}

// type PriceHistory struct {
// 	ID         int
// 	MenuItemId int
// 	OldPrice   float64
// 	NewPrice   float64
// 	CreatedAt  time.Time
// }
