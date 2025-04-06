// mapper/menu.go
package mapper

import (
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/models"
)

func ToOrderModel(e entity.Order) models.Order {
	return models.Order{
		ID:                  int(e.ID),
		CustomerName:        e.CustomerName,
		TotalAmount:         e.TotalAmount,
		Status:              e.Status.String(),
		PaymentMethod:       e.PaymentMethod.String(),
		SpecialInstructions: models.JSONB(e.SpecialInstructions),
		CreatedAt:           e.CreatedAt,
		UpdatedAt:           e.UpdatedAt,
	}
}

func ToOrderEntity(m models.Order, items []entity.OrderItem) entity.Order {
	return entity.Order{
		ID:                  int64(m.ID),
		CustomerName:        m.CustomerName,
		TotalAmount:         m.TotalAmount,
		Status:              entity.ParseStatus(m.Status),
		PaymentMethod:       entity.ParsePaymentMethod(m.PaymentMethod),
		SpecialInstructions: entity.JSONB(m.SpecialInstructions),
		OrderItems:          items,
		CreatedAt:           m.CreatedAt,
		UpdatedAt:           m.UpdatedAt,
	}
}
