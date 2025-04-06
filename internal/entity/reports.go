package entity

type PopularItem struct {
	ProductId   int    `json:"product_id"`
	ProductName string `json:"product_name"`
	Sold        int    `json:"total_quantity"`
}

type TotalItemsByPeriod struct {
	Period       string           `json:"period"`
	Month        string           `json:"month,omitempty"`
	Year         int              `json:"year,omitempty"`
	OrderedItems []map[string]int `json:"ordered_items"`
}

type SearchResult struct {
	MenuItems []SearchMenuItem `json:"menu_items,omitempty"`
	Orders    []SearchOrder    `json:"orders,omitempty"`
	Matches   int              `json:"total_matches"`
}

type SearchMenuItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string
	Price       float64 `json:"price"`
	Relevance   float64 `json:"relevance"`
}

type SearchOrder struct {
	ID           int      `json:"id"`
	CustomerName string   `json:"customer_name"`
	Items        []string `json:"items"`
	Total        float64  `json:"total"`
	Relevance    float64  `json:"relevance"`
}

type NumberOfOrderedItemsByPeriod struct {
	OrderedItems map[string]int `json:"items"`
}
