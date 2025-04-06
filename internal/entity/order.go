package entity

import (
	"time"
)

type Order struct {
	ID                  int64
	CustomerName        string
	TotalAmount         float64
	Status              OrderStatus
	PaymentMethod       PaymentMethod
	SpecialInstructions JSONB
	OrderItems          []OrderItem
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type OrderStatus int

type OrderItem struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}

const (
	OrderPending OrderStatus = iota
	OrderProcessing
	OrderCompleted
	OrderCancelled
)

func ParseStatus(s string) OrderStatus {
	switch s {
	case "pending":
		return OrderPending
	case "processing":
		return OrderProcessing
	case "completed":
		return OrderCompleted
	case "cancelled":
		return OrderCancelled
	default:
		return OrderStatus(-1)
	}
}

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

func ParsePaymentMethod(s string) PaymentMethod {
	switch s {
	case "cash":
		return PaymentCash
	case "card":
		return PaymentCard
	case "online":
		return PaymentOnline
	default:
		return PaymentMethod(-1)
	}
}

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
