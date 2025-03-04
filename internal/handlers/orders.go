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
}

type OrderHandler struct {
	Service OrderService
	Logger  *slog.Logger
}

func NewOrderHandler(orderService OrderService, logger *slog.Logger) *OrderHandler {
	return &OrderHandler{orderService, logger}
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
