package handlers

import (
	"context"
	"frappuccino-alem/internal/entity"
	"log/slog"
	"net/http"
)

type ReportService interface {
	GetPopularItems(ctx context.Context) ([]entity.PopularItem, error)
	GetTotalSales(ctx context.Context) (float64, error)
	// GetFilterSearch ??
}

type ReportHandler struct {
	service ReportService
	logger  *slog.Logger
}

func NewReportHandler(service ReportService, logger *slog.Logger) *ReportHandler {
	return &ReportHandler{service, logger}
}

func (h *ReportHandler) RegisterEndpoints(mux *http.ServeMux) {
	mux.HandleFunc("GET /reports/total-sales", h.GetTotalSales)
	mux.HandleFunc("GET /reports/total-sales/", h.GetTotalSales)

	mux.HandleFunc("GET /reports/popular-items", h.GetPopularItems)
	mux.HandleFunc("GET /reports/popular-items/", h.GetPopularItems)

	mux.HandleFunc("GET /reports/search", h.GetFilterSearch)
	mux.HandleFunc("GET /reports/orderedItemsByPeriod", h.GetTotalItemsByPeriod)
}

func (h *ReportHandler) GetTotalSales(w http.ResponseWriter, r *http.Request) {
}

func (h *ReportHandler) GetPopularItems(w http.ResponseWriter, r *http.Request) {
}

// Search through orders, menu items, and customers with partial matching and ranking
func (h *ReportHandler) GetFilterSearch(w http.ResponseWriter, r *http.Request) {
	// url := r.URL.Query()

	// querystring := url.Get("q")
	// filter := url.Get("filter")
	// minPrice := url.Get("minPrice")
	// maxPrice := url.Get("maxPrice")
}

func (h *ReportHandler) GetTotalItemsByPeriod(w http.ResponseWriter, r *http.Request) {
	// url := r.URL.Query()

	// period:=url.Get("period")
	// month:=url.Get("month")
	// year:=url.Get("year")
}
