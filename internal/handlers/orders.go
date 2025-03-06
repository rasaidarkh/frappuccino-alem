package handlers

import (
	"context"
	"frappuccino-alem/models"
	"log/slog"
	"net/http"
)

type OrderService interface {
	CreateOrder(ctx context.Context, item models.Order) (string, error)
	GetAllOrders(ctx context.Context) ([]models.Order, error)
	GetOrderById(ctx context.Context, OrderId string) (models.Order, error)
	UpdateOrderById(ctx context.Context, OrderId string, item models.Order) error
	DeleteOrderById(ctx context.Context, OrderId string) error
	CloseOrderById(ctx context.Context, OrderId string) error
	GetNumberOfOrderedItems(ctx context.Context, startDate, endDate string) (map[string]int, error)
}

type OrderHandler struct {
	service OrderService
	logger  *slog.Logger
}

func NewOrderHandler(service OrderService, logger *slog.Logger) *OrderHandler {
	return &OrderHandler{service, logger}
}

func (h *OrderHandler) RegisterEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("POST /orders", h.CreateOrder)
	mux.HandleFunc("POST /orders/", h.CreateOrder)

	mux.HandleFunc("GET /orders", h.GetAllOrders)
	mux.HandleFunc("GET /orders/", h.GetAllOrders)

	mux.HandleFunc("GET /orders/{id}", h.GetOrderById)
	mux.HandleFunc("GET /orders/{id}/", h.GetOrderById)

	mux.HandleFunc("PUT /orders/{id}", h.UpdateOrderById)
	mux.HandleFunc("PUT /orders/{id}/", h.UpdateOrderById)

	mux.HandleFunc("DELETE /orders/{id}", h.DeleteOrderById)
	mux.HandleFunc("DELETE /orders/{id}/", h.DeleteOrderById)

	mux.HandleFunc("POST /orders/{id}/close", h.CloseOrderById)
	mux.HandleFunc("POST /orders/{id}/close/", h.CloseOrderById)

	mux.HandleFunc("GET /orders/numberOfOrderedItems", h.GetNumberOfOrderedItems)
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) GetAllOrders(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) GetOrderById(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) UpdateOrderById(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) DeleteOrderById(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) CloseOrderById(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) GetNumberOfOrderedItems(w http.ResponseWriter, r *http.Request) {
}
