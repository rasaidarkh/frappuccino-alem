package handlers

import (
	"context"
	"frappuccino-alem/models"
	"log/slog"
	"net/http"
)

type MenuService interface {
	CreateMenuItem(ctx context.Context, item models.MenuItem) (string, error)
	GetAllMenuItems(ctx context.Context) ([]models.MenuItem, error)
	GetMenuItemById(ctx context.Context, id string) (models.MenuItem, error)
	UpdateMenuItemById(ctx context.Context, id string, item models.MenuItem) error
	DeleteMenuItemById(ctx context.Context, id string) error
}

type MenuHandler struct {
	Service MenuService
	Logger  *slog.Logger
}

func NewMenuHandler(service MenuService, logger *slog.Logger) *MenuHandler {
	return &MenuHandler{service, logger}
}

func (h *MenuHandler) RegisterEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("POST /menu", h.createMenuItem)
	mux.HandleFunc("POST /menu/", h.createMenuItem)

	mux.HandleFunc("GET /menu", h.getAllMenuItems)
	mux.HandleFunc("GET /menu/", h.getAllMenuItems)

	mux.HandleFunc("GET /menu/{id}", h.getMenuItemById)
	mux.HandleFunc("GET /menu/{id}/", h.getMenuItemById)

	mux.HandleFunc("PUT /menu/{id}", h.updateMenuItemById)
	mux.HandleFunc("PUT /menu/{id}/", h.updateMenuItemById)

	mux.HandleFunc("DELETE /menu/{id}", h.deleteMenuItemById)
	mux.HandleFunc("DELETE /menu/{id}/", h.deleteMenuItemById)
}

func (h *MenuHandler) createMenuItem(w http.ResponseWriter, r *http.Request) {
}

func (h *MenuHandler) getAllMenuItems(w http.ResponseWriter, r *http.Request) {
}

func (h *MenuHandler) getMenuItemById(w http.ResponseWriter, r *http.Request) {
}

func (h *MenuHandler) updateMenuItemById(w http.ResponseWriter, r *http.Request) {
}

func (h *MenuHandler) deleteMenuItemById(w http.ResponseWriter, r *http.Request) {
}

func validateMenuItem(item models.MenuItem) error {
	return nil
}
