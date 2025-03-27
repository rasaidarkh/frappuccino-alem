package store

import (
	"context"
	"database/sql"
	"frappuccino-alem/models"
)

type OrderStore struct {
	db *sql.DB
}

func NewOrderStore(db *sql.DB) *OrderStore {
	return &OrderStore{db}
}
func (r *OrderStore) CreateOrder(ctx context.Context, order models.Order) (string, error) {
	const op = "Store.CreateOrder"

	// logic here ...

	return "", nil
}

func (r *OrderStore) GetAllOrders(ctx context.Context) ([]models.Order, error) {
	const op = "Store.GetAllOrders"
	var orders []models.Order

	// logic here ...

	return orders, nil
}

func (r *OrderStore) GetOrderById(ctx context.Context, orderId string) (models.Order, error) {
	const op = "Store.GetOrderById"
	var order models.Order

	// logic here ...

	return order, nil
}

func (r *OrderStore) UpdateOrderById(ctx context.Context, orderId string, order models.Order) error {
	const op = "Store.UpdateOrderById"

	// logic here ...

	return nil
}

func (r *OrderStore) DeleteOrderById(ctx context.Context, orderId string) error {
	const op = "Store.DeleteOrderById"

	// logic here ...

	return nil
}

func (r *OrderStore) CloseOrderById(ctx context.Context, orderId string) error {
	const op = "Store.CloseOrderById"

	// logic here ...

	return nil
}

func (r *OrderStore) GetNumberOfOrderedItems(ctx context.Context, startDate, endDate string) (map[string]int, error) {
	const op = "Store.GetNumberOfOrderedItems"

	// logic here ...

	return nil, nil
}
