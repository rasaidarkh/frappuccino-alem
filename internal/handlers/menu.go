package handlers

import (
	"context"
	"errors"
	"fmt"
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/types"
	"frappuccino-alem/internal/store"
	"frappuccino-alem/internal/utils"
	"log/slog"
	"net/http"
)

type MenuService interface {
	CreateMenuItem(ctx context.Context, item entity.MenuItem) (string, error)
	GetPaginatedMenuItems(ctx context.Context, pagination *types.Pagination) (*types.PaginationResponse[entity.MenuItem], error)
	GetMenuItemById(ctx context.Context, id string) (entity.MenuItem, error)
	UpdateMenuItemById(ctx context.Context, id string, item entity.MenuItem) error
	DeleteMenuItemById(ctx context.Context, id string) error
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
	var dto types.MenuItemCreateRequest
	if err := utils.ParseJSON(r, &dto); err != nil {
		h.logger.Error("Failed to parse request body", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// add validation
	entityItem := dto.ToEntity()
	id, err := h.service.CreateMenuItem(r.Context(), entityItem)
	if err != nil {
		h.logger.Error("Failed to create menu item", "error", err.Error())
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, map[string]string{"id": id})
}

func (h *MenuHandler) getPaginatedMenuItems(w http.ResponseWriter, r *http.Request) {
	pagination, err := types.NewPaginationFromRequest(r, []types.SortOption{
		types.SortByID,
		types.SortByName,
		types.SortByPrice,
		types.SortByCreatedAt,
		types.SortByUpdatedAt,
	})
	if err != nil {
		h.logger.Error("Failed to parse menu item pagination request", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	response, err := h.service.GetPaginatedMenuItems(r.Context(), pagination)
	if err != nil {
		h.logger.Error("Failed to get paginated menu items", slog.Any("pagination", pagination), "error", err.Error())
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	h.logger.Info("Succeded to get menu items page")
	utils.WriteJSON(w, http.StatusOK, response)
}

func (h *MenuHandler) getMenuItemById(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing menu item ID"))
		return
	}

	item, err := h.service.GetMenuItemById(r.Context(), id)
	if err != nil {
		if errors.Is(err, store.ErrNotFound) {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("menu item not found"))
			return
		}
		h.logger.Error("Failed to get menu item", "id", id, "error", err.Error())
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, item)
}

func (h *MenuHandler) updateMenuItemById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing menu item ID"))
		return
	}
	var dto types.MenuItemUpdateRequest
	if err := utils.ParseJSON(r, &dto); err != nil {
		h.logger.Error("Failed to parse request body", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if dto.Name == nil && dto.Description == nil && dto.Price == nil &&
		dto.Categories == nil && dto.Allergens == nil &&
		dto.Metadata == nil && dto.Ingredients == nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("at least one field must be provided"))
		return
	}

	entityItem := dto.ToEntity()
	if err := h.service.UpdateMenuItemById(r.Context(), id, entityItem); err != nil {
		h.logger.Error("Failed to update menu item", "id", id, "error", err.Error())
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *MenuHandler) deleteMenuItemById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("missing menu item ID"))
		return
	}

	if err := h.service.DeleteMenuItemById(r.Context(), id); err != nil {
		if errors.Is(err, store.ErrNotFound) {
			utils.WriteError(w, http.StatusNotFound, fmt.Errorf("menu item not found"))
			return
		}
		h.logger.Error("Failed to delete menu item", "id", id, "error", err.Error())
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
