package service

import (
	"context"
	"fmt"
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/dto"
	"frappuccino-alem/internal/store"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, item entity.Order) (int64, error)
	GetAllOrders(ctx context.Context, pagination *dto.Pagination) ([]entity.Order, error)
	GetTotalOrdersCount(ctx context.Context) (int, error)
	GetOrderById(ctx context.Context, OrderId string) (entity.Order, error)
	UpdateOrderById(ctx context.Context, OrderId string, item entity.Order) error
	DeleteOrderById(ctx context.Context, OrderId string) error
	CloseOrderById(ctx context.Context, OrderId string) error
	GetNumberOfOrderedItems(ctx context.Context, startDate, endDate string) (map[string]int, error)
}

type OrderService struct {
	inventoryRepo store.InventoryRepository
	menuRepo      store.MenuRepository
	orderRepo     OrderRepository
}

func NewOrderService(inventoryRepo store.InventoryRepository, menuRepo store.MenuRepository, orderRepo OrderRepository) *OrderService {
	return &OrderService{
		inventoryRepo,
		menuRepo,
		orderRepo,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, order entity.Order) (entity.Order, error) {
	const op = "service.CreateOrder"
	// Validate order items
	usedIngredients := make(map[int64]float64)

	for i, item := range order.OrderItems {
		menuItem, err := s.menuRepo.GetMenuItemById(ctx, item.ID)
		if err != nil {
			return order, err
		}
		for _, ing := range menuItem.Ingredients {
			// check if inventory quantity of existingItem is enough
			storedIng, err := s.inventoryRepo.GetInventoryItemById(ctx, ing.ID)
			if err != nil {
				return order, err
			}
			if storedIng.Quantity < ing.Quantity*float64(item.Quantity) {
				return order, fmt.Errorf("%s: not enough inventory for ingredient %s", op, ing.ItemName)
			}
			usedIngredients[ing.ID] += ing.Quantity * float64(item.Quantity)
			if storedIng.Quantity < usedIngredients[ing.ID] {
				return order, fmt.Errorf("%s: not enough inventory for ingredient %s", op, ing.ItemName)
			}
		}
		order.OrderItems[i].Name = menuItem.Name
		order.OrderItems[i].Price = menuItem.Price
		order.TotalAmount += menuItem.Price * float64(item.Quantity)
		fmt.Printf("%#v %#v\n", order.TotalAmount, item)
	}
	order.Status = entity.OrderPending
	orderID, err := s.orderRepo.CreateOrder(ctx, order)
	if err != nil {
		return order, fmt.Errorf("%s: failed to create order, %w", op, err)
	}
	order.ID = orderID
	return order, nil
}

func (s *OrderService) GetPaginatedOrders(ctx context.Context, pagination *dto.Pagination) (*dto.PaginationResponse[entity.Order], error) {
	const op = "service.GetPaginatedOrders"

	totalItems, err := s.orderRepo.GetTotalOrdersCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	totalPages := (totalItems + pagination.PageSize - 1) / pagination.PageSize

	orders, err := s.orderRepo.GetAllOrders(ctx, pagination)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &dto.PaginationResponse[entity.Order]{
		CurrentPage: pagination.Page,
		HasNextPage: pagination.Page < totalPages,
		PageSize:    pagination.PageSize,
		TotalPages:  totalPages,
		Data:        orders,
	}, nil
}

func (s *OrderService) GetOrderById(ctx context.Context, orderId string) (entity.Order, error) {
	const op = "service.GetOrderById"
	// logic here ...
	order, err := s.orderRepo.GetOrderById(ctx, orderId)
	if err != nil {
		return entity.Order{}, fmt.Errorf("%s: %w", op, err)
	}

	return order, nil
}

func (s *OrderService) UpdateOrderById(ctx context.Context, orderId string, order entity.Order) error {
	const op = "service.UpdateOrderById"
	// logic here ...
	err := s.orderRepo.UpdateOrderById(ctx, orderId, order)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *OrderService) DeleteOrderById(ctx context.Context, orderId string) error {
	const op = "service.DeleteOrderById"
	// logic here ...
	err := s.orderRepo.DeleteOrderById(ctx, orderId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *OrderService) CloseOrderById(ctx context.Context, orderId string) error {
	const op = "service.CloseOrderById"
	// logic here ...
	err := s.orderRepo.CloseOrderById(ctx, orderId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *OrderService) GetNumberOfOrderedItems(ctx context.Context, startDate, endDate string) (map[string]int, error) {
	const op = "service.GetNumberOfOrderedItems"
	//logic here ...
	OrderMap, err := s.orderRepo.GetNumberOfOrderedItems(ctx, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return OrderMap, nil
}
