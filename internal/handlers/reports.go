package handlers

import (
	"context"
	"errors"
	"fmt"
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/utils"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ReportService interface {
	GetTotalSales(ctx context.Context) (float64, error)
	GetPopularItems(ctx context.Context) ([]entity.PopularItem, error)
	GetFilterSearch(ctx context.Context, search string, filter string, minPrice float64, maxPrice float64) (entity.SearchResult, error)
	GetTotalItemsByPeriod(ctx context.Context, period string, month int, year int) (entity.TotalItemsByPeriod, error)
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

func (h *ReportHandler) GetPopularItems(w http.ResponseWriter, r *http.Request) {
	popularItems, err := h.service.GetPopularItems(r.Context())
	if err != nil {
		h.logger.Error("failed to get popular items", "error", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if popularItems == nil {
		h.logger.Info("no popular items found")
		http.Error(w, "No popular items found", http.StatusNotFound)
	}

	utils.WriteJSON(w, http.StatusOK, popularItems)
}

func (h *ReportHandler) GetTotalSales(w http.ResponseWriter, r *http.Request) {
	totalSales, err := h.service.GetTotalSales(r.Context())
	if err != nil {
		h.logger.Error("failed to get total sales", "error", err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]float64{
		"total_sales": totalSales})
}

func (h *ReportHandler) GetFilterSearch(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query()

	querystring := url.Get("q")
	filter := url.Get("filter")
	minPriceStr := url.Get("minPrice")
	maxPriceStr := url.Get("maxPrice")

	if querystring == "" {
		utils.WriteError(w, http.StatusBadRequest, errors.New("query parameter is required"))
		return
	}

	if minPriceStr == "" {
		minPriceStr = "0"
	}
	if maxPriceStr == "" {
		maxPriceStr = "0"
	}

	minPriceFloat, err := strconv.ParseFloat(minPriceStr, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid minPrice parameter"))
		return
	}

	maxPriceFloat, err := strconv.ParseFloat(maxPriceStr, 64)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid maxPrice parameter"))
		return
	}

	if minPriceFloat < 0 {
		utils.WriteError(w, http.StatusBadRequest, errors.New("minPrice cannot be negative"))
		return
	}
	if maxPriceFloat < 0 {
		utils.WriteError(w, http.StatusBadRequest, errors.New("maxPrice cannot be negative"))
		return
	}
	if minPriceFloat > maxPriceFloat {
		utils.WriteError(w, http.StatusBadRequest, errors.New("minPrice cannot be greater than maxPrice"))
		return
	}

	data, err := h.service.GetFilterSearch(r.Context(), querystring, filter, minPriceFloat, maxPriceFloat)
	if err != nil {
		h.logger.Error("could not get filter search", "error", err.Error())
		utils.WriteError(w, http.StatusInternalServerError, errors.New("could not get filter search"))
		return
	}

	utils.WriteJSON(w, http.StatusOK, data)
}

func (h *ReportHandler) GetTotalItemsByPeriod(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	period := q.Get("period")
	if period == "" {
		utils.WriteError(w, http.StatusBadRequest, errors.New("period parameter is required"))
		return
	}

	switch period {
	case "day":
		monthStr := q.Get("month")
		yearStr := q.Get("year")

		if monthStr == "" {
			utils.WriteError(w, http.StatusBadRequest, errors.New("month parameter is required for period=day"))
			return
		}

		month, err := monthNameToNumber(monthStr)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid month: %v", err))
			return
		}

		year, err := parseYearParam(yearStr)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, errors.New("valid year parameter required"))
			return
		}

		data, err := h.service.GetTotalItemsByPeriod(r.Context(), "day", month, year)
		if err != nil {
			h.logger.Error("could not get total items by period", "error", err.Error())
			utils.WriteError(w, http.StatusInternalServerError, errors.New("could not get total items by period"))
			return
		}

		utils.WriteJSON(w, http.StatusOK, data)

	case "month":
		yearStr := q.Get("year")
		year, err := parseYearParam(yearStr)
		if err != nil {
			utils.WriteError(w, http.StatusBadRequest, errors.New("valid year parameter required"))
			return
		}

		data, err := h.service.GetTotalItemsByPeriod(r.Context(), "month", 0, year)
		if err != nil {
			h.logger.Error("could not get total items by period", "error", err.Error())
			utils.WriteError(w, http.StatusInternalServerError, errors.New("could not get total items by period"))
			return
		}

		utils.WriteJSON(w, http.StatusOK, data)

	default:
		utils.WriteError(w, http.StatusBadRequest, errors.New("invalid period value"))
	}
}

func parseYearParam(yearStr string) (int, error) {
	if yearStr == "" {
		return time.Now().Year(), nil
	}
	return strconv.Atoi(yearStr)
}

func monthNameToNumber(month string) (int, error) {
	month = strings.ToLower(strings.TrimSpace(month))
	months := map[string]int{
		"january":   1,
		"february":  2,
		"march":     3,
		"april":     4,
		"may":       5,
		"june":      6,
		"july":      7,
		"august":    8,
		"september": 9,
		"october":   10,
		"november":  11,
		"december":  12,
	}
	if m, ok := months[month]; ok {
		return m, nil
	}
	return 0, fmt.Errorf("invalid month: %s", month)
}
