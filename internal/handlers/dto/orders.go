package dto

import (
	"fmt"
)

type OrderRequest struct {
	CustomerName        *string             `json:"customer_name"`
	PaymentMethod       *string             `json:"payment_method"`
	OrderItems          *[]OrderItemRequest `json:"order_items"`
	SpecialInstructions *string             `json:"special_instructions"`
}

type OrderItemRequest struct {
	ItemID   *int64   `json:"item_id"`
	Quantity *float64 `json:"quantity"`
}

func (r OrderRequest) Validate() error {
	if r.CustomerName == nil || *r.CustomerName == "" {
		return fmt.Errorf("customer_name is required")
	}
	if r.PaymentMethod == nil || *r.PaymentMethod == "" {
		return fmt.Errorf("payment_method is required")
	}
	if r.OrderItems == nil || len(*r.OrderItems) == 0 {
		return fmt.Errorf("order_items are required")
	}
	for _, item := range *r.OrderItems {
		if item.ItemID == nil || *item.ItemID <= 0 {
			return fmt.Errorf("item_id must be greater than 0")
		}
		if item.Quantity == nil || *item.Quantity <= 0 {
			return fmt.Errorf("quantity must be greater than 0")
		}
	}
	return nil
}

// func (r OrderRequest) MapToEntity() *entity.Order {
// 	var paymentMethod entity.PaymentMethod
// 	switch *r.PaymentMethod {
// 	case "cash":
// 		paymentMethod = entity.PaymentCash
// 	case "card":
// 		paymentMethod = entity.PaymentCard
// 	case "online":
// 		paymentMethod = entity.PaymentOnline
// 	default:
// 		paymentMethod = -1
// 	}

// 	order := &entity.Order{
// 		CustomerName:        *r.CustomerName,
// 		PaymentMethod:       paymentMethod,
// 		SpecialInstructions: r.SpecialInstructions,
// 	}

// 	if r.OrderItems != nil {
// 		orderItems := make([]entity.MenuItem, len(*r.OrderItems))
// 		for i, item := range *r.OrderItems {
// 			orderItems[i] = entity.MenuItem{
// 				ID:   *item.ItemID,
// 				Quantity: *item.Quantity,
// 			}
// 		}
// 		order.OrderItems = orderItems
// 	}

// 	return order
// }
