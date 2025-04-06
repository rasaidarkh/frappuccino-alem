package store

import (
	"context"
	"database/sql"
	"fmt"
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/dto"
	"frappuccino-alem/models"
	"frappuccino-alem/models/mapper"
	"strings"
)

type OrderStore struct {
	db *sql.DB
}

func NewOrderStore(db *sql.DB) *OrderStore {
	return &OrderStore{db}
}
func (r *OrderStore) CreateOrder(ctx context.Context, order entity.Order) (int64, error) {
	const op = "Store.CreateOrder"
	var id int64

	err := runInTx(r.db, func(tx *sql.Tx) error {
		modelOrder := mapper.ToOrderModel(order)

		err := tx.QueryRowContext(ctx,
			`INSERT INTO orders (customer_name, status, total_amount, payment_method, special_instructions)
			VALUES ($1, $2, $3, $4, $5)
			RETURNING id`,
			modelOrder.CustomerName, modelOrder.Status, modelOrder.TotalAmount, modelOrder.PaymentMethod, modelOrder.SpecialInstructions,
		).Scan(&id)
		if err != nil {
			return fmt.Errorf("insert order: %w", err)
		}

		if len(order.OrderItems) > 0 {
			valueStrings := make([]string, 0, len(order.OrderItems))
			valueArgs := make([]interface{}, 0, len(order.OrderItems)*4)

			for i, item := range order.OrderItems {
				valueStrings = append(valueStrings, fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))
				valueArgs = append(valueArgs, id, item.ID, item.Quantity, item.Price*float64(item.Quantity))
			}

			_, err = tx.ExecContext(ctx,
				fmt.Sprintf(`INSERT INTO order_items (order_id, menu_item_id, quantity, price_at_order)
					VALUES %s`, strings.Join(valueStrings, ",")),
				valueArgs...)
			if err != nil {
				return fmt.Errorf("insert order items: %w", err)
			}
		}

		return nil
	})
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (r *OrderStore) GetAllOrders(ctx context.Context, pagination *dto.Pagination) ([]entity.Order, error) {
	const op = "Store.GetAllOrders"

	query := `
		SELECT id, customer_name, payment_method, total_amount, status, created_at, updated_at
		FROM orders o
	`

	if pagination.SortBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", pagination.SortBy)
	}

	offset := (pagination.Page - 1) * pagination.PageSize
	query += fmt.Sprintf(" LIMIT %d OFFSET %d", pagination.PageSize, offset)

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var modelItems []models.Order
	for rows.Next() {
		var model models.Order
		err := rows.Scan(
			&model.ID,
			&model.CustomerName,
			&model.PaymentMethod,
			&model.TotalAmount,
			&model.Status,
			&model.CreatedAt,
			&model.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		modelItems = append(modelItems, model)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	entities := make([]entity.Order, 0, len(modelItems))
	for _, model := range modelItems {
		items, err := r.getMenuItemsForOrder(ctx, model.ID)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		entities = append(entities, mapper.ToOrderEntity(model, items))
	}

	return entities, nil
}

func (r *OrderStore) getMenuItemsForOrder(ctx context.Context, orderID int) ([]entity.OrderItem, error) {
	const op = "Store.getMenuItemsForOrder"
	query := `
        SELECT
            mi.id,
            mi.name,
            mi.price,
            oi.quantity
        FROM order_items oi
        JOIN menu_items mi ON mi.id = oi.menu_item_id
        WHERE oi.order_id = $1
    `

	rows, err := r.db.QueryContext(ctx, query, orderID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var items []entity.OrderItem
	for rows.Next() {
		var item entity.OrderItem
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Price,
			&item.Quantity,
		)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		items = append(items, item)
	}

	return items, nil
}

func (r *OrderStore) GetOrderById(ctx context.Context, orderId string) (entity.Order, error) {
	const op = "Store.GetOrderById"
	var order entity.Order

	// logic here ...

	return order, nil
}

func (r *OrderStore) GetTotalOrdersCount(ctx context.Context) (int, error) {
	const op = "Store.GetTotalOrdersCount"

	var total int
	err := r.db.QueryRowContext(ctx, "SELECT COUNT(*) FROM orders").Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return total, nil
}

func (r *OrderStore) UpdateOrderById(ctx context.Context, orderId string, order entity.Order) error {
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
