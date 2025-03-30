package models

import (
	"time"

	"github.com/lib/pq"
)

type MenuItem struct {
	ID          int            `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Categories  pq.StringArray `json:"categories"`
	Allergens   pq.StringArray `json:"allergens"`
	Metadata    JSONB          `json:"metadata"`
	CreatedAt   time.Time      `json:"created_at"`
}

type MenuItemIngredient struct {
	ID               int64   `json:"id"`
	MenuItemID       int64   `json:"menu_item_id"`
	InventoryID      int64   `json:"inventory_id"`
	QuantityRequired float64 `json:"quantity_required"` // Amount needed per recipe
}

type PriceHistory struct {
	ID         int       `json:"id"`
	MenuItemId int       `json:"menu_item_id"`
	OldPrice   float64   `json:"old_price"`
	NewPrice   float64   `json:"new_price"`
	CreatedAt  time.Time `json:"created_at"`
}
