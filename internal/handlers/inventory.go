package handlers

import (
	"context"
	"errors"
	"fmt"
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/types"
	"frappuccino-alem/internal/utils"
	"log/slog"
	"net/http"
	"strconv"
)

type InventoryService interface {
	CreateInventoryItem(ctx context.Context, item entity.InventoryItem) (int64, error)
	GetPaginatedInventoryItems(ctx context.Context, pagination *types.Pagination) (*types.PaginationResponse[entity.InventoryItem], error)
	GetInventoryItemById(ctx context.Context, InventoryId int64) (entity.InventoryItem, error)
	DeleteInventoryItemById(ctx context.Context, InventoryId int64) (entity.InventoryItem, error)
	UpdateInventoryItemById(ctx context.Context, InventoryId int64, item types.InventoryItemRequest) error
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

	mux.HandleFunc("GET /inventory", h.getPaginatedInventoryItems)
	mux.HandleFunc("GET /inventory/", h.getPaginatedInventoryItems)

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
		h.logger.Error("Failed to parse inventory item request", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid request payload"))
		return
	}

	if err := validateInventoryItem(item); err != nil {
		h.logger.Error("Some of the fields are incorrect", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, errors.New("Validation failed"))
		return
	}

	entity := item.MapToInventoryItemEntity()

	id, err := h.service.CreateInventoryItem(r.Context(), entity)
	if err != nil {
		h.logger.Error("Failed to create new inventory item", "error", err.Error())
		utils.WriteError(w, http.StatusInternalServerError, errors.New("Failed to create new inventory item"))
		return
	}
	h.logger.Info("Succeeded to create new inventory item", slog.Int64("id", id))
	utils.WriteMessage(w, http.StatusCreated, "Created new inventory item")
}

func (h *InventoryHandler) getPaginatedInventoryItems(w http.ResponseWriter, r *http.Request) {
	pagination, err := types.NewPaginationFromRequest(r, []types.SortOption{
		types.SortByID,
		types.SortByName,
		types.SortByQuantity,
		types.SortByCreatedAt,
		types.SortByUpdatedAt,
	})
	if err != nil {
		h.logger.Error("Failed to parse inventory item pagination request", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.service.GetPaginatedInventoryItems(r.Context(), pagination)
	if err != nil {
		h.logger.Error("Failed to get paginated inventory items", slog.Any("pagination", pagination), "error", err.Error())
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	h.logger.Info("Succeded to get inventory items page")
	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *InventoryHandler) getInventoryItemById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		h.logger.Error("Cannot convert inventory id to integer value", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, errors.New("Cannot convert inventory id to integer value"))
		return
	}
	item, err := h.service.GetInventoryItemById(r.Context(), int64(id))
	if err != nil {
		h.logger.Error("Failed to get inventory item", slog.Int("id", id), "error", err.Error())
		utils.WriteError(w, http.StatusNotFound, errors.New(fmt.Sprintf("Item with id %v not found", id)))
		return
	}
	h.logger.Info("Succeded to get inventory item - ", slog.Int("id", item.ID), slog.String("Name", item.Name))
	utils.WriteJSON(w, http.StatusOK, item)
}

func (h *InventoryHandler) updateInventoryItemById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		h.logger.Error("Cannot convert inventory id to integer value", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, errors.New("Cannot convert inventory id to integer value"))
		return
	}
	var itemRequest types.InventoryItemRequest
	if err := utils.ParseJSON(r, &itemRequest); err != nil {
		h.logger.Error("Failed to parse inventory item request", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, errors.New("Invalid request payload"))
		return
	}
	h.logger.Debug("itemRequest", itemRequest)
	err = h.service.UpdateInventoryItemById(r.Context(), int64(id), itemRequest)
	if err != nil {
		h.logger.Error("Failed to update inventory item", slog.Int("id", id), "error", err.Error())
		utils.WriteError(w, http.StatusInternalServerError, errors.New("Failed to update inventory item"))
		return
	}
	h.logger.Info("Succeeded to update inventory item", slog.Int("id", id))
	utils.WriteMessage(w, http.StatusOK, "Updated inventory item")
}

func (h *InventoryHandler) deleteInventoryItemById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		h.logger.Error("Cannot convert inventory id to integer value", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, errors.New("Cannot convert inventory id to integer value"))
		return
	}
	item, err := h.service.DeleteInventoryItemById(r.Context(), int64(id))
	if err != nil {
		h.logger.Error("Failed to get inventory item", slog.Int("id", id), "error", err.Error())
		utils.WriteError(w, http.StatusNotFound, errors.New(fmt.Sprintf("Item with id %v not found", id)))
		return
	}
	h.logger.Info("Succeded to delete inventory item", slog.Int("id", item.ID), slog.String("Name", item.Name))
	utils.WriteMessage(w, http.StatusNotFound, fmt.Sprintf("Deleted inventory item %v", item.Name))
}

func (h *InventoryHandler) GetLeftOvers(w http.ResponseWriter, r *http.Request) {
	validSortByOptions := []types.SortOption{
		types.SortByQuantity,
		types.SortByPrice,
	}

	pagination, err := types.NewPaginationFromRequest(r, validSortByOptions)
	if err != nil {
		h.logger.Error("Invalid query parameters", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid query parameters"))
		return
	}

	response, err := h.service.GetPaginatedInventoryItems(r.Context(), pagination)
	if err != nil {
		h.logger.Error("Failed to get leftovers", "error", err.Error())
		utils.WriteError(w, http.StatusInternalServerError, errors.New("failed to retrieve leftover items"))
		return
	}

	h.logger.Info("Successfully retrieved leftover items")
	utils.WriteJSON(w, http.StatusOK, response)
}

func validateInventoryItem(item types.InventoryItemRequest) error {
	if item.Name == nil {
		return errors.New("item name is required")
	}
	if item.Quantity == nil {
		return errors.New("item quantity is required")
	}
	if item.UnitType == nil {
		return errors.New("item unit is required")
	}

	if *item.Name == "" {
		return errors.New("item name cannot be empty")
	}
	if *item.Quantity <= 0 {
		return errors.New("stock level must be greater than zero")
	}
	if *item.UnitType == "" {
		return errors.New("unit type cannot be empty")
	}

	if item.Price == nil {
		return errors.New("item price is required")
	}
	if *item.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	return nil
}
