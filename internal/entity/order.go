package entity

import (
	"time"
)

type Order struct {
	ID                  int
	CustomerName        string
	TotalAmount         float64
	Status              OrderStatus
	PaymentMethod       PaymentMethod
	SpecialInstructions JSONB
	CreatedAt           time.Time
	UpdatedAt           time.Time
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
