package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"frappuccino-alem/internal/handlers/dto"
	"frappuccino-alem/internal/service"
	"frappuccino-alem/internal/store"
	"frappuccino-alem/internal/utils"
)

type MenuHandler struct {
	service service.MenuService
	logger  *slog.Logger
}

func NewMenuHandler(service service.MenuService, logger *slog.Logger) *MenuHandler {
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
	var req dto.MenuItemRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		h.logError("Failed to parse request body", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Failed to parse request body"))
		return
	}
	if err := req.Validate(); err != nil {
		h.logError("Invalid request", err)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("Invalid request: %v", err))
		return
	}

	entityItem := req.MapToEntity()
	item, err := h.service.CreateMenuItem(r.Context(), entityItem)
	if err != nil {
		h.logError("Failed to create menu item", err)
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("Failed to create menu item: %v", err))
		return
	}
	utils.WriteJSON(w, http.StatusCreated, dto.MenuItemToResponse(item))
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

	paginatedData, err := h.service.GetPaginatedMenuItems(r.Context(), pagination)
	if err != nil {
		h.logError("Failed to get menu items", err)
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	response := dto.PaginationResponse[dto.MenuItemResponse]{
		CurrentPage: paginatedData.CurrentPage,
		HasNextPage: paginatedData.HasNextPage,
		PageSize:    paginatedData.PageSize,
		TotalPages:  paginatedData.TotalPages,
	}
	for _, item := range paginatedData.Data {
		response.Data = append(response.Data, dto.MenuItemToResponse(item))
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

	utils.WriteJSON(w, http.StatusOK, dto.MenuItemToDetailedResponse(entityItem))
}

func (h *MenuHandler) updateMenuItemById(w http.ResponseWriter, r *http.Request) {
	id, err := parsePathID(r, "id")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	var req dto.MenuItemRequest
	if err := utils.ParseJSON(r, &req); err != nil {
		h.logger.Error("Failed to parse inventory item request", "error", err.Error())
		utils.WriteError(w, http.StatusBadRequest, errors.New("Invalid request payload"))
		return
	}

	if req.Ingredients != nil {
		for _, ing := range *req.Ingredients {
			if *ing.Quantity <= 0 {
				utils.WriteError(w, http.StatusBadRequest,
					fmt.Errorf("quantity_used must be greater than 0"))
				return
			}
		}
	}

	h.logger.Debug("update request ", "menuRequest", req)
	err = h.service.UpdateMenuItemById(r.Context(), int64(id), req)
	if err != nil {
		h.logger.Error("Failed to update menu item", slog.Int64("id", id), "error", err.Error())
		utils.WriteError(w, http.StatusInternalServerError, errors.New("Failed to update menu item"))
		return
	}
	h.logger.Info("Succeeded to update menu item", slog.Int64("id", id))
	utils.WriteMessage(w, http.StatusOK, "Updated menu item")
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
