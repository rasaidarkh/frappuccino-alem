package store

import (
	"context"
	"database/sql"
	"fmt"
	"frappuccino-alem/internal/entity"

	"github.com/lib/pq"
)

type ReportStore struct {
	db *sql.DB
}

func NewReportStore(db *sql.DB) *ReportStore {
	return &ReportStore{db}
}

func (r *ReportStore) GetPopularItems(ctx context.Context) ([]entity.PopularItem, error) {
	const op = "Store.GetPopularItems"

	var popularItems []entity.PopularItem

	query := `
		SELECT 
			order_items.menu_item_id, 
			menu_items.name AS item, 
			COUNT(*) AS order_count
		FROM order_items
		JOIN orders ON order_items.order_id = orders.id
		JOIN menu_items ON order_items.menu_item_id = menu_items.id
		WHERE orders.status = 'completed'
		GROUP BY order_items.menu_item_id, menu_items.name
		ORDER BY order_count DESC
		LIMIT 10
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()
	for rows.Next() {
		var item entity.PopularItem
		if err := rows.Scan(&item.ProductId, &item.ProductName, &item.Sold); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		popularItems = append(popularItems, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return popularItems, nil
}

func (r *ReportStore) GetTotalSales(ctx context.Context) (float64, error) {
	const op = "Store.GetPopularItems"

	var totalSales float64
	query := `SELECT COALESCE(SUM(total_amount), 0) FROM orders WHERE status = 'completed'`
	if err := r.db.QueryRowContext(ctx, query).Scan(&totalSales); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return totalSales, nil
}

func (s *ReportStore) SearchMenuItems(ctx context.Context, query string, minPrice, maxPrice float64) ([]entity.SearchMenuItem, error) {
	const op = "ReportStore.SearchMenuItems"

	q := `
         SELECT id, name, description, price,
               ts_rank(search_vector, websearch_to_tsquery('english', $1)) AS relevance
        FROM menu_items
        WHERE search_vector @@ websearch_to_tsquery('english', $1)
        AND (CAST($2 AS DECIMAL) = 0 OR price >= CAST($2 AS DECIMAL))
        AND (CAST($3 AS DECIMAL) = 0 OR price <= CAST($3 AS DECIMAL))
        ORDER BY relevance DESC
    `

	rows, err := s.db.QueryContext(ctx, q,
		query,
		minPrice,
		maxPrice,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var items []entity.SearchMenuItem
	for rows.Next() {
		var item entity.SearchMenuItem
		if err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Description,
			&item.Price,
			&item.Relevance,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		items = append(items, item)
	}

	return items, nil
}

func (s *ReportStore) SearchOrders(ctx context.Context, query string, minPrice, maxPrice float64) ([]entity.SearchOrder, error) {
	const op = "ReportStore.SearchOrders"

	q := `
	SELECT 
		o.id, 
		o.customer_name, 
		o.total_amount,
		ts_rank(o.search_vector, websearch_to_tsquery('english', $1)) AS relevance,
		array_agg(mi.name) AS items
	FROM orders o
	JOIN order_items oi ON o.id = oi.order_id
	JOIN menu_items mi ON oi.menu_item_id = mi.id
	WHERE 
		o.search_vector @@ websearch_to_tsquery('english', $1)
		AND (CAST($2 AS DECIMAL) = 0 OR o.total_amount >= CAST($2 AS DECIMAL))
		AND (CAST($3 AS DECIMAL) = 0 OR o.total_amount <= CAST($3 AS DECIMAL))
	GROUP BY o.id
	ORDER BY relevance DESC
	`

	rows, err := s.db.QueryContext(ctx, q,
		query,
		minPrice,
		maxPrice,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var orders []entity.SearchOrder
	for rows.Next() {
		var order entity.SearchOrder
		var items pq.StringArray
		if err := rows.Scan(
			&order.ID,
			&order.CustomerName,
			&order.Total,
			&order.Relevance,
			&items,
		); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		order.Items = items
		orders = append(orders, order)
	}

	return orders, nil
}

func (s *ReportStore) GetTotalItemsByDay(ctx context.Context, month int, year int) (map[int]int, error) {
	const op = "ReportStore.GetTotalItemsByDay"

	query := `
        SELECT EXTRACT(DAY FROM created_at)::integer AS day, 
               COUNT(id) AS order_count
        FROM orders
        WHERE EXTRACT(MONTH FROM created_at) = $1 
          AND EXTRACT(YEAR FROM created_at) = $2
        GROUP BY day
        ORDER BY day
    `

	rows, err := s.db.QueryContext(ctx, query, month, year)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	results := make(map[int]int)
	for rows.Next() {
		var day int
		var count int
		if err := rows.Scan(&day, &count); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		results[day] = count
	}

	return results, nil
}

func (s *ReportStore) GetTotalItemsByMonth(ctx context.Context, year int) (map[int]int, error) {
	const op = "ReportStore.GetTotalItemsByMonth"

	query := `
        SELECT EXTRACT(MONTH FROM created_at)::integer AS month, 
               COUNT(id) AS order_count
        FROM orders
        WHERE EXTRACT(YEAR FROM created_at) = $1
        GROUP BY month
        ORDER BY month
    `

	rows, err := s.db.QueryContext(ctx, query, year)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	results := make(map[int]int)
	for rows.Next() {
		var month int
		var count int
		if err := rows.Scan(&month, &count); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		results[month] = count
	}

	return results, nil
}
