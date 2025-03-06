package repository

import (
	"context"
	"database/sql"
	"frappuccino-alem/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db}
}
func (r *OrderRepository) CreateOrder(ctx context.Context, order models.Order) (string, error) {
	const op = "repository.CreateOrder"

	// logic here ...

	return "", nil
}

func (r *OrderRepository) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	const op = "repository.GetAllOrders"
	var orders []models.Order

	// logic here ...

	return orders, nil
}

func (r *OrderRepository) GetOrderById(ctx context.Context, orderId string) (models.Order, error) {
	const op = "repository.GetOrderById"
	var order models.Order

	// logic here ...

	return order, nil
}

func (r *OrderRepository) UpdateOrderById(ctx context.Context, orderId string, order models.Order) error {
	const op = "repository.UpdateOrderById"

	// logic here ...

	return nil
}

func (r *OrderRepository) DeleteOrderById(ctx context.Context, orderId string) error {
	const op = "repository.DeleteOrderById"

	// logic here ...

	return nil
}

func (r *OrderRepository) CloseOrderById(ctx context.Context, orderId string) error {
	const op = "repository.CloseOrderById"

	// logic here ...

	return nil
}

func (r *OrderRepository) GetNumberOfOrderedItems(ctx context.Context, startDate, endDate string) (map[string]int, error) {
	const op = "repository.GetNumberOfOrderedItems"

	// logic here ...

	return nil, nil
}
