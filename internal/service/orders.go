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
	GetNumberOfOrderedItems(ctx context.Context, startDate, endDate string) (map[string]int, error)
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(Repo OrderRepository) *OrderService {
	return &OrderService{Repo}
}

func (s *OrderService) CreateOrder(ctx context.Context, order models.Order) (string, error) {
	const op = "service.CreateOrder"
	// logic here ...
	orderID, err := s.repo.CreateOrder(ctx, order)
	if err != nil {
		return "", fmt.Errorf("%s: failed to create order, %w", op, err)
	}

	return orderID, nil
}

func (s *OrderService) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	const op = "service.GetAllOrders"
	// logic here ...
	orders, err := s.repo.GetAllOrders(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return orders, nil
}

func (s *OrderService) GetOrderById(ctx context.Context, orderId string) (models.Order, error) {
	const op = "service.GetOrderById"
	// logic here ...
	order, err := s.repo.GetOrderById(ctx, orderId)
	if err != nil {
		return models.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	return order, nil
}

func (s *OrderService) UpdateOrderById(ctx context.Context, orderId string, order models.Order) error {
	const op = "service.UpdateOrderById"
	// logic here ...
	err := s.repo.UpdateOrderById(ctx, orderId, order)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *OrderService) DeleteOrderById(ctx context.Context, orderId string) error {
	const op = "service.DeleteOrderById"
	// logic here ...
	err := s.repo.DeleteOrderById(ctx, orderId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *OrderService) CloseOrderById(ctx context.Context, orderId string) error {
	const op = "service.CloseOrderById"
	// logic here ...
	err := s.repo.CloseOrderById(ctx, orderId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *OrderService) GetNumberOfOrderedItems(ctx context.Context, startDate, endDate string) (map[string]int, error) {
	const op = "service.GetNumberOfOrderedItems"
	//logic here ...
	OrderMap, err := s.repo.GetNumberOfOrderedItems(ctx, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return OrderMap, nil
}
