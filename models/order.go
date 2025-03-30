package models

import (
	"time"
)

type Order struct {
	ID                  int       `json:"id"`
	CustomerName        string    `json:"customer_name"`
	TotalAmount         float64   `json:"total_amount"`
	Status              string    `json:"status"`
	PaymentMethod       string    `json:"payment_method"`
	SpecialInstructions JSONB     `json:"special_instructions,omitempty"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at,omitempty"`
}

type OrderStatusHistory struct {
	ID             int       `json:"id"`
	OrderID        int       `json:"order_id"`
	PreviousStatus string    `json:"previous_status"`
	NewStatus      string    `json:"new_status"`
	ChangedAt      time.Time `json:"changed_at"`
}

type OrderItem struct {
	ID            int64   `json:"id"`
	OrderID       int64   `json:"order_id"`
	MenuItemID    int64   `json:"menu_item_id"`
	Quantity      int     `json:"quantity"`
	PriceAtTime   float64 `json:"price_at_time"`
	Customization JSONB   `json:"customization,omitempty"`
}
