package dto

import (
	"fmt"
	"frappuccino-alem/internal/entity"
	"time"
)

type OrderRequest struct {
	CustomerName        *string             `json:"customer_name"`
	PaymentMethod       *string             `json:"payment_method"`
	SpecialInstructions *entity.JSONB       `json:"special_instructions"`
	Items               *[]OrderItemRequest `json:"menu_items"`
}

type OrderItemRequest struct {
	MenuItemID int64 `json:"id"`
	Quantity   int64 `json:"quantity"`
}

type OrderResponse struct {
	ID                  int64               `json:"id"`
	CustomerName        string              `json:"customer_name"`
	Status              string              `json:"status"`
	TotalAmount         float64             `json:"total_amount"`
	PaymentMethod       string              `json:"payment_method"`
	SpecialInstructions entity.JSONB        `json:"special_instructions"`
	Items               []OrderItemResponse `json:"items"`
	CreatedAt           time.Time           `json:"created_at"`
}

type OrderItemResponse struct {
	MenuItemName string  `json:"menu_item_name"`
	Quantity     int     `json:"quantity"`
	UnitPrice    float64 `json:"unit_price"`
	TotalPrice   float64 `json:"total_price"`
}

func (r OrderRequest) Validate() error {
	if r.CustomerName == nil || *r.CustomerName == "" {
		return fmt.Errorf("customer_name is required")
	}
	if r.PaymentMethod == nil || *r.PaymentMethod == "" {
		return fmt.Errorf("payment_method is required")
	}
	if r.Items == nil || len(*r.Items) == 0 {
		return fmt.Errorf("menu_items are required")
	}
	if r.PaymentMethod == nil || !entity.ParsePaymentMethod(*r.PaymentMethod).IsValid() {
		return fmt.Errorf("invalid payment_method %v", entity.ParsePaymentMethod(*r.PaymentMethod).IsValid())
	}
	for _, item := range *r.Items {
		if item.MenuItemID <= 0 {
			return fmt.Errorf("menu_item id must be greater than 0")
		}
		if item.Quantity <= 0 {
			return fmt.Errorf("quantity must be greater than 0")
		}
	}
	return nil
}

func (r OrderRequest) MapToEntity() entity.Order {
	var paymentMethod entity.PaymentMethod
	switch *r.PaymentMethod {
	case "cash":
		paymentMethod = entity.PaymentCash
	case "card":
		paymentMethod = entity.PaymentCard
	case "online":
		paymentMethod = entity.PaymentOnline
	default:
		paymentMethod = -1
	}

	order := entity.Order{
		PaymentMethod: paymentMethod,
	}

	if r.CustomerName != nil {
		order.CustomerName = *r.CustomerName
	}
	if r.SpecialInstructions != nil {
		order.SpecialInstructions = *r.SpecialInstructions
	}

	if r.Items != nil {
		orderItems := make([]entity.OrderItem, len(*r.Items))
		for i, item := range *r.Items {
			orderItems[i] = entity.OrderItem{
				ID:       item.MenuItemID,
				Quantity: item.Quantity,
			}
		}
		order.OrderItems = orderItems
	}

	return order
}

func OrderToResponse(entity entity.Order) OrderResponse {
	orderItems := make([]OrderItemResponse, 0)
	for _, i := range entity.OrderItems {
		orderItems = append(orderItems, OrderItemResponse{
			MenuItemName: i.Name,
			Quantity:     int(i.Quantity),
			UnitPrice:    i.Price,
			TotalPrice:   i.Price * float64(i.Quantity),
		})
	}
	return OrderResponse{
		ID:                  entity.ID,
		CustomerName:        entity.CustomerName,
		Status:              entity.Status.String(),
		TotalAmount:         entity.TotalAmount,
		PaymentMethod:       entity.PaymentMethod.String(),
		SpecialInstructions: entity.SpecialInstructions,
		Items:               orderItems,
		CreatedAt:           entity.CreatedAt,
	}
}
