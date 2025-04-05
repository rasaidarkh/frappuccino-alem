package handlers

import (
	"context"
	"errors"
	"fmt"
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/dto"
	"frappuccino-alem/internal/store"
	"frappuccino-alem/internal/utils"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type MenuService interface {
	CreateMenuItem(ctx context.Context, item entity.MenuItem) (int64, error)
	GetPaginatedMenuItems(ctx context.Context, pagination *dto.Pagination) (*dto.PaginationResponse[entity.MenuItem], error)
	GetMenuItemById(ctx context.Context, id int64) (entity.MenuItem, error)
	UpdateMenuItemById(ctx context.Context, id int64, item entity.MenuItem) error
	DeleteMenuItemById(ctx context.Context, id int64) error
}

type MenuHandler struct {
	service MenuService
	logger  *slog.Logger
}

func NewMenuHandler(service MenuService, logger *slog.Logger) *MenuHandler {
	return &MenuHandler{service, logger}
}

func (h *MenuHandler) RegisterEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("POST /menu", h.createMenuItem)
	mux.HandleFunc("POST /menu/", h.createMenuItem)

	mux.HandleFunc("GET /menu", h.getPaginatedMenuItems)
	mux.HandleFunc("GET /menu/", h.getPaginatedMenuItems)

	mux.HandleFunc("GET /menu/{id}", h.getMenuItemById)
	mux.HandleFunc("GET /menu/{id}/", h.getMenuItemById)

	mux.HandleFunc("PUT /menu/{id}", h.updateMenuItemById)
	mux.HandleFunc("PUT /menu/{id}/", h.updateMenuItemById)

	mux.HandleFunc("DELETE /menu/{id}", h.deleteMenuItemById)
	mux.HandleFunc("DELETE /menu/{id}/", h.deleteMenuItemById)
}

func (h *MenuHandler) createMenuItem(w http.ResponseWriter, r *http.Request) {
	var req dto.MenuItemCreateRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		h.logError("Failed to parse request body", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := validateMenuItemIngredients(req.Ingredients); err != nil {
		h.logError("Invalid ingredients", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	entityItem := req.ToEntity()
	id, err := h.service.CreateMenuItem(r.Context(), entityItem)
	if err != nil {
		h.logError("Failed to create menu item", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, dto.CreatedResponse{ID: id})
}

func (h *MenuHandler) getPaginatedMenuItems(w http.ResponseWriter, r *http.Request) {
	pagination, err := dto.NewPaginationFromRequest(r, []dto.SortOption{
		dto.SortByID,
		dto.SortByName,
		dto.SortByPrice,
		dto.SortByCreatedAt,
		dto.SortByUpdatedAt,
	})
	if err != nil {
		h.logError("Invalid pagination request", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.service.GetPaginatedMenuItems(r.Context(), pagination)
	if err != nil {
		h.logError("Failed to get menu items", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *MenuHandler) getMenuItemById(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	entityItem, err := h.service.GetMenuItemById(r.Context(), id)
	if err != nil {
		h.handleNotFoundOrError(w, "menu item", id, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, dto.ToMenuItemResponse(entityItem))
}

func (h *MenuHandler) updateMenuItemById(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	var req dto.MenuItemUpdateRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		h.logError("Failed to parse request body", err)
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if req.IsEmpty() {
		utils.WriteError(w, http.StatusBadRequest, errors.New("at least one field must be provided"))
		return
	}

	if req.Ingredients != nil {
		for _, ing := range *req.Ingredients {
			if ing.QuantityUsed <= 0 {
				utils.WriteError(w, http.StatusBadRequest,
					fmt.Errorf("quantity_used must be greater than 0"))
				return
			}
		}
	}

	if err := h.service.UpdateMenuItemById(r.Context(), id, req.ToEntity()); err != nil {
		h.handleNotFoundOrError(w, "menu item", id, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MenuHandler) deleteMenuItemById(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.service.DeleteMenuItemById(r.Context(), id); err != nil {
		h.handleNotFoundOrError(w, "menu item", id, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Helper functions
func (h *MenuHandler) logError(message string, err error) {
	h.logger.Error(message, "error", err.Error())
}

func (h *MenuHandler) handleNotFoundOrError(w http.ResponseWriter, resourceType string, id int64, err error) {
	if errors.Is(err, store.ErrNotFound) {
		utils.WriteError(w, http.StatusNotFound, fmt.Errorf("%s with ID %d not found", resourceType, id))
		return
	}
	h.logError(fmt.Sprintf("Failed to process %s", resourceType), err)
	utils.WriteError(w, http.StatusInternalServerError, err)
}

func parsePathID(r *http.Request, param string) (int64, error) {
	idStr := r.PathValue(param)
	if idStr == "" {
		return 0, fmt.Errorf("missing %s ID", param)
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid ID format: must be integer")
	}
	return id, nil
}

func validateMenuItemIngredients(ingredients []dto.MenuItemIngredientDTO) error {
	for _, ing := range ingredients {
		if ing.QuantityUsed <= 0 {
			return fmt.Errorf("invalid quantity_used: must be greater than 0")
		}
	}
	return nil
}

func validateMenuItem(m entity.MenuItem) error {
	if strings.TrimSpace(m.Name) == "" {
		return errors.New("menu item name cannot be empty")
	}
	if strings.TrimSpace(m.Description) == "" {
		return errors.New("menu item description cannot be empty")
	}
	if m.Price <= 0 {
		return errors.New("menu item price must be greater than zero")
	}
	if len(m.Categories) == 0 {
		return errors.New("at least one category is required")
	}
	if len(m.Ingredients) == 0 {
		return errors.New("menu item must have at least one ingredient")
	}
	for i, ing := range m.Ingredients {
		if ing.InventoryID <= 0 {
			return errors.New("ingredient at index " + strconv.Itoa(i) + " has invalid inventory ID")
		}
		if ing.QuantityUsed <= 0 {
			return errors.New("ingredient at index " + strconv.Itoa(i) + " must have positive quantity used")
		}
	}
	return nil
}
