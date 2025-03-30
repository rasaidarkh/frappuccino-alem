package handlers

import (
	"context"
	"errors"
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/types"
	"frappuccino-alem/internal/utils"
	"log/slog"
	"net/http"
)

type InventoryService interface {
	CreateInventoryItem(ctx context.Context, item entity.InventoryItem) (int64, error)
	GetAllInventoryItems(ctx context.Context) ([]entity.InventoryItem, error)
	GetInventoryItemById(ctx context.Context, InventoryId int64) (entity.InventoryItem, error)
	DeleteInventoryItemById(ctx context.Context, InventoryId int64) error
	UpdateInventoryItemById(ctx context.Context, InventoryId int64, item entity.InventoryItem) error
}

type InventoryHandler struct {
	service InventoryService
	logger  *slog.Logger
}

func NewInventoryHandler(service InventoryService, logger *slog.Logger) *InventoryHandler {
	return &InventoryHandler{service, logger}
}

func (h *InventoryHandler) RegisterEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("POST /inventory", h.createInventoryItem)
	mux.HandleFunc("POST /inventory/", h.createInventoryItem)

	mux.HandleFunc("GET /inventory", h.getAllInventoryItems)
	mux.HandleFunc("GET /inventory/", h.getAllInventoryItems)

	mux.HandleFunc("GET /inventory/{id}", h.getInventoryItemById)
	mux.HandleFunc("GET /inventory/{id}/", h.getInventoryItemById)

	mux.HandleFunc("PUT /inventory/{id}", h.updateInventoryItemById)
	mux.HandleFunc("PUT /inventory/{id}/", h.updateInventoryItemById)

	mux.HandleFunc("DELETE /inventory/{id}", h.deleteInventoryItemById)
	mux.HandleFunc("DELETE /inventory/{id}/", h.deleteInventoryItemById)

	mux.HandleFunc("GET /inventory/getLeftOvers", h.GetLeftOvers)
}

func (h *InventoryHandler) createInventoryItem(w http.ResponseWriter, r *http.Request) {
	var item types.InventoryItemRequest
	if err := utils.ParseJSON(r, &item); err != nil {
		h.logger.Error("Failed to parse inventory item request", "error", err)
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid request payload"))
		return
	}

	if err := validateInventoryItem(item); err != nil {
		h.logger.Error("Some of the fields are incorrect", "error", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	entity := item.MapToInventoryItemEntity()

	id, err := h.service.CreateInventoryItem(r.Context(), entity)
	if err != nil {
		h.logger.Error("Failed to create new inventory item", "error", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	h.logger.Info("Succeeded to create new inventory item", slog.Int64("id", id))
	utils.WriteJSON(w, http.StatusCreated, "created new inventory item")
}

func (h *InventoryHandler) getAllInventoryItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.service.GetAllInventoryItems(r.Context())
	if err != nil {
		h.logger.Error("Failed to get all inventory items", "error", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	h.logger.Info("Succeeded to get all inventory items")
	utils.WriteJSON(w, http.StatusOK, items)
}

func (h *InventoryHandler) getInventoryItemById(w http.ResponseWriter, r *http.Request) {
}

func (h *InventoryHandler) updateInventoryItemById(w http.ResponseWriter, r *http.Request) {
}

func (h *InventoryHandler) deleteInventoryItemById(w http.ResponseWriter, r *http.Request) {
}

func (h *InventoryHandler) GetLeftOvers(w http.ResponseWriter, r *http.Request) {
}

func validateInventoryItem(item types.InventoryItemRequest) error {
	if item.Name == "" {
		return errors.New("item  cannot be empty")
	}
	if item.Quantity <= 0 {
		return errors.New("stock level must be greater than zero")
	}
	if item.UnitType == "" {
		return errors.New("unit type cannot be empty")
	}
	return nil
}
