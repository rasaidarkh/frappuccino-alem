package handlers

import (
	"context"
	"errors"
	entity "frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/utils"
	"log/slog"
	"net/http"
)

type InventoryService interface {
	CreateInventoryItem(ctx context.Context, item entity.InventoryItem) (string, error)
	GetAllInventoryItems(ctx context.Context) ([]entity.InventoryItem, error)
	GetInventoryItemById(ctx context.Context, InventoryId string) (entity.InventoryItem, error)
	DeleteInventoryItemById(ctx context.Context, InventoryId string) error
	UpdateInventoryItemById(ctx context.Context, InventoryId string, item entity.InventoryItem) error
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
	var item entity.InventoryItem

	if err := utils.ParseJSON(r, &item); err != nil {
		h.logger.Error("Failed to parse inventory item request", "error", err)
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid request payload"))
		return
	}

	if err := validateInventoryItem(item); err != nil {
		h.logger.Warn("Wrong  inventory item body", "error", err)
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid request payload"))
		return
	}

}

func (h *InventoryHandler) getAllInventoryItems(w http.ResponseWriter, r *http.Request) {
}

func (h *InventoryHandler) getInventoryItemById(w http.ResponseWriter, r *http.Request) {
}

func (h *InventoryHandler) updateInventoryItemById(w http.ResponseWriter, r *http.Request) {
}

func (h *InventoryHandler) deleteInventoryItemById(w http.ResponseWriter, r *http.Request) {
}

func (h *InventoryHandler) GetLeftOvers(w http.ResponseWriter, r *http.Request) {
}

func validateInventoryItem(item entity.InventoryItem) error {
	if item.Name == "" {
		return errors.New("item  cannot be empty")
	}
	if item.StockLevel <= 0 {
		return errors.New("stock level must be greater than zero")
	}
	if item.UnitType == "" {
		return errors.New("unit type cannot be empty")
	}
	return nil
}
