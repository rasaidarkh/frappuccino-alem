package handlers

import (
	"context"
	"frappuccino-alem/models"
	"log/slog"
	"net/http"
)

type InventoryService interface {
	CreateInventoryItem(ctx context.Context, item models.InventoryItem) (string, error)
	GetAllInventoryItems(ctx context.Context) ([]models.InventoryItem, error)
	GetInventoryItemById(ctx context.Context, InventoryId string) (models.InventoryItem, error)
	DeleteInventoryItemById(ctx context.Context, InventoryId string) error
	UpdateInventoryItemById(ctx context.Context, InventoryId string, item models.InventoryItem) error
}

type InventoryHandler struct {
	Service InventoryService
	Logger  *slog.Logger
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
}

func (h *InventoryHandler) createInventoryItem(w http.ResponseWriter, r *http.Request) {

}

func (h *InventoryHandler) getAllInventoryItems(w http.ResponseWriter, r *http.Request) {

}

func (h *InventoryHandler) getInventoryItemById(w http.ResponseWriter, r *http.Request) {

}

func (h *InventoryHandler) updateInventoryItemById(w http.ResponseWriter, r *http.Request) {

}

func (h *InventoryHandler) deleteInventoryItemById(w http.ResponseWriter, r *http.Request) {

}

func validateInventoryItem(item models.InventoryItem) error {
	return nil
}
