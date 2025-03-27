package models

import (
	"database/sql"
	"time"
)

type Order struct {
	ID                  int           `json:"id"`
	CustomerName        string        `json:"customer_name"`
	TotalAmount         float64       `json:"total_amount"`
	Status              OrderStatus   `json:"status"`
	PaymentMethod       PaymentMethod `json:"payment_method"`
	SpecialInstructions JSONB         `json:"special_instructions,omitempty"`
	CreatedAt           time.Time     `json:"created_at"`
	UpdatedAt           sql.NullTime  `json:"updated_at,omitempty"`
}

type OrderStatusHistory struct {
	ID             int         `json:"id"`
	OrderID        int         `json:"order_id"`
	PreviousStatus OrderStatus `json:"previous_status"`
	NewStatus      OrderStatus `json:"new_status"`
	ChangedAt      time.Time   `json:"changed_at"`
}

type OrderItem struct {
	ID            int64   `json:"id"`
	OrderID       int64   `json:"order_id"`
	MenuItemID    int64   `json:"menu_item_id"`
	Quantity      int     `json:"quantity"`
	PriceAtTime   float64 `json:"price_at_time"`
	Customization JSONB   `json:"customization,omitempty"`
}

type OrderStatus int

const (
	OrderPending OrderStatus = iota
	OrderProcessing
	OrderCompleted
	OrderCancelled
)

func (s OrderStatus) String() string {
	switch s {
	case OrderPending:
		return "pending"
	case OrderProcessing:
		return "processing"
	case OrderCompleted:
		return "completed"
	case OrderCancelled:
		return "cancelled"
	default:
		return "unknown"
	}
}

func (s OrderStatus) IsValid() bool {
	switch s {
	case OrderPending, OrderProcessing, OrderCompleted, OrderCancelled:
		return true
	}
	return false
}

// -------------------------------------------------------------------------

type PaymentMethod int

const (
	PaymentCash PaymentMethod = iota
	PaymentCard
	PaymentOnline
)

func (m PaymentMethod) String() string {
	switch m {
	case PaymentCash:
		return "cash"
	case PaymentCard:
		return "card"
	case PaymentOnline:
		return "online"
	default:
		return "unknown"
	}
}

func (p PaymentMethod) IsValid() bool {
	switch p {
	case PaymentCash, PaymentCard, PaymentOnline:
		return true
	}
	return false
}
