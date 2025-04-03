package handlers

import (
	"context"
	"frappuccino-alem/internal/handlers/types"
	"frappuccino-alem/internal/utils"
	"frappuccino-alem/models"
	"log/slog"
	"net/http"
)

type MenuService interface {
	CreateMenuItem(ctx context.Context, item models.MenuItem) (string, error)
	GetPaginatedMenuItems(ctx context.Context, pagination *types.Pagination) (*types.PaginationResponse[models.MenuItem], error)
	GetMenuItemById(ctx context.Context, id string) (models.MenuItem, error)
	UpdateMenuItemById(ctx context.Context, id string, item models.MenuItem) error
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
}

func (h *MenuHandler) updateMenuItemById(w http.ResponseWriter, r *http.Request) {
}

func (h *MenuHandler) deleteMenuItemById(w http.ResponseWriter, r *http.Request) {
}
