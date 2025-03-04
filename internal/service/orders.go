package service

import (
	"context"
	"fmt"
	"frappuccino-alem/models"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, item models.Order) (string, error)
	GetAllOrders(ctx context.Context) ([]models.Order, error)
	GetOrderById(ctx context.Context, OrderId string) (models.Order, error)
	UpdateOrderById(ctx context.Context, OrderId string, item models.Order) error
	DeleteOrderById(ctx context.Context, OrderId string) error
	CloseOrderById(ctx context.Context, OrderId string) error
}

type OrderService struct {
	OrderRepo OrderRepository
}

func NewOrderService(OrderRepo OrderRepository) *OrderService {
	return &OrderService{OrderRepo}
}

func (s *OrderService) CreateOrder(ctx context.Context, order models.Order) (string, error) {
	const op = "service.CreateOrder"
	// logic here ...
	orderID, err := s.OrderRepo.CreateOrder(ctx, order)
	if err != nil {
		return "", fmt.Errorf("%s: failed to create order, %w", op, err)
	}

	return orderID, nil
}

func (s *OrderService) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	const op = "service.GetAllOrders"
	// logic here ...
	orders, err := s.OrderRepo.GetAllOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return orders, nil
}

func (s *OrderService) GetOrderById(ctx context.Context, orderId string) (models.Order, error) {
	const op = "service.GetOrderById"
	// logic here ...
	order, err := s.OrderRepo.GetOrderById(ctx, orderId)
	if err != nil {
		return models.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	return order, nil
}

func (s *OrderService) UpdateOrderById(ctx context.Context, orderId string, order models.Order) error {
	const op = "service.UpdateOrderById"
	// logic here ...
	err := s.OrderRepo.UpdateOrderById(ctx, orderId, order)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *OrderService) DeleteOrderById(ctx context.Context, orderId string) error {
	const op = "service.DeleteOrderById"
	// logic here ...
	err := s.OrderRepo.DeleteOrderById(ctx, orderId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *OrderService) CloseOrderById(ctx context.Context, orderId string) error {
	const op = "service.CloseOrderById"
	// logic here ...
	err := s.OrderRepo.CloseOrderById(ctx, orderId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
