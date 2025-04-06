package handlers

import (
	"context"
	"fmt"
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/dto"
	"frappuccino-alem/internal/utils"
	"log/slog"
	"net/http"
)

type OrderService interface {
	CreateOrder(ctx context.Context, item entity.Order) (entity.Order, error)
	GetPaginatedOrders(ctx context.Context, pagination *dto.Pagination) (*dto.PaginationResponse[entity.Order], error)
	GetOrderById(ctx context.Context, OrderId string) (entity.Order, error)
	UpdateOrderById(ctx context.Context, OrderId string, item entity.Order) error
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
	mux.HandleFunc("POST /orders", h.createOrder)
	mux.HandleFunc("POST /orders/", h.createOrder)

	mux.HandleFunc("GET /orders", h.getPaginatedOrders)
	mux.HandleFunc("GET /orders/", h.getPaginatedOrders)

	mux.HandleFunc("GET /orders/{id}", h.getOrderById)
	mux.HandleFunc("GET /orders/{id}/", h.getOrderById)

	mux.HandleFunc("PUT /orders/{id}", h.updateOrderById)
	mux.HandleFunc("PUT /orders/{id}/", h.updateOrderById)

	mux.HandleFunc("DELETE /orders/{id}", h.deleteOrderById)
	mux.HandleFunc("DELETE /orders/{id}/", h.deleteOrderById)

	mux.HandleFunc("POST /orders/{id}/close", h.closeOrderById)
	mux.HandleFunc("POST /orders/{id}/close/", h.closeOrderById)

	mux.HandleFunc("GET /orders/numberOfOrderedItems", h.getNumberOfOrderedItems)
	mux.HandleFunc("GET /orders/numberOfOrderedItems/", h.getNumberOfOrderedItems)
}

func (h *OrderHandler) createOrder(w http.ResponseWriter, r *http.Request) {
	var req dto.OrderRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		h.logger.Error("Failed to parse order request", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request payload"))
		return
	}
	if err := req.Validate(); err != nil {
		h.logger.Error("Invalid order request", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid request payload: %v", err))
		return
	}

	entityItem := req.MapToEntity()
	item, err := h.service.CreateOrder(r.Context(), entityItem)
	if err != nil {
		h.logger.Error("Failed to create menu item", "error", err.Error())
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to create menu item: %v", err))
		return
	}
	utils.WriteJSON(w, http.StatusCreated, dto.OrderToResponse(item))

}

func (h *OrderHandler) getPaginatedOrders(w http.ResponseWriter, r *http.Request) {
	pagination, err := dto.NewPaginationFromRequest(r, []dto.SortOption{
		dto.SortByID,
		dto.SortByName,
		dto.SortByPrice,
		dto.SortByCreatedAt,
		dto.SortByUpdatedAt,
	})
	if err != nil {
		h.logger.Error("Invalid pagination request", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	paginatedData, err := h.service.GetPaginatedOrders(r.Context(), pagination)
	if err != nil {
		h.logger.Error("Failed to get orders", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := dto.PaginationResponse[dto.OrderResponse]{
		CurrentPage: paginatedData.CurrentPage,
		HasNextPage: paginatedData.HasNextPage,
		PageSize:    paginatedData.PageSize,
		TotalPages:  paginatedData.TotalPages,
	}
	for _, item := range paginatedData.Data {
		response.Data = append(response.Data, dto.OrderToResponse(item))
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *OrderHandler) getOrderById(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) updateOrderById(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) deleteOrderById(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) closeOrderById(w http.ResponseWriter, r *http.Request) {
}

func (h *OrderHandler) getNumberOfOrderedItems(w http.ResponseWriter, r *http.Request) {
}
