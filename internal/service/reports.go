package service

import (
	"context"
	"fmt"
	"frappuccino-alem/internal/entity"
	"strconv"
	"strings"
	"time"
)

type ReportRepository interface {
	GetPopularItems(ctx context.Context) ([]entity.PopularItem, error)
	GetTotalSales(ctx context.Context) (float64, error)
	SearchMenuItems(ctx context.Context, query string, minPrice, maxPrice float64) ([]entity.SearchMenuItem, error)
	SearchOrders(ctx context.Context, query string, minPrice, maxPrice float64) ([]entity.SearchOrder, error)
	GetTotalItemsByDay(ctx context.Context, month int, year int) (map[int]int, error)
	GetTotalItemsByMonth(ctx context.Context, year int) (map[int]int, error)
}

type ReportService struct {
	repo ReportRepository
}

func NewReportService(repo ReportRepository) *ReportService {
	return &ReportService{repo}
}

func (s *ReportService) GetPopularItems(ctx context.Context) ([]entity.PopularItem, error) {
	const op = "service.GetPopularItems"
	// logic here ...

	popularItems, err := s.repo.GetPopularItems(ctx)
	if err != nil {
		return []entity.PopularItem{}, fmt.Errorf("%s: %w", op, err)
	}

	return popularItems, nil
}

func (s *ReportService) GetTotalSales(ctx context.Context) (float64, error) {
	const op = "service.GetTotalSales"
	// logic here ...
	totalSales, err := s.repo.GetTotalSales(ctx)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}
	return totalSales, nil
}

func (s *ReportService) GetTotalItemsByPeriod(ctx context.Context, period string, month int, year int) (entity.TotalItemsByPeriod, error) {
	const op = "service.GetTotalItemsByPeriod"

	result := entity.TotalItemsByPeriod{}

	switch period {
	case "day":
		dayCounts, err := s.repo.GetTotalItemsByDay(ctx, month, year)
		if err != nil {
			return result, fmt.Errorf("%s: %w", op, err)
		}

		result.Period = "day"
		result.Month = time.Month(month).String()
		result.Year = year
		result.OrderedItems = make([]map[string]int, 0)

		for day, count := range dayCounts {
			result.OrderedItems = append(result.OrderedItems,
				map[string]int{strconv.Itoa(day): count})
		}

	case "month":
		monthCounts, err := s.repo.GetTotalItemsByMonth(ctx, year)
		if err != nil {
			return result, fmt.Errorf("%s: %w", op, err)
		}

		result.Period = "month"
		result.Year = year
		result.OrderedItems = make([]map[string]int, 0)

		for monthNum, count := range monthCounts {
			monthName := strings.ToLower(time.Month(monthNum).String())
			result.OrderedItems = append(result.OrderedItems,
				map[string]int{monthName: count})
		}

	default:
		return result, fmt.Errorf("%s: invalid period", op)
	}

	return result, nil
}

func (s *ReportService) GetFilterSearch(ctx context.Context, search string, filter string, minPrice float64, maxPrice float64) (entity.SearchResult, error) {
	const op = "service.GetFilterSearch"
	var results entity.SearchResult
	if filter != "menu" && filter != "orders" && filter != "" {
		return results, fmt.Errorf("%s: invalid filter", op)
	}
	var menuResults []entity.SearchMenuItem
	var orderResults []entity.SearchOrder
	var err error
	if filter == "menu" {
		menuResults, err = s.repo.SearchMenuItems(ctx, search, minPrice, maxPrice)
		if err != nil {
			return entity.SearchResult{}, fmt.Errorf("%s: %w", op, err)
		}
	}
	if filter == "orders" {
		orderResults, err = s.repo.SearchOrders(ctx, search, minPrice, maxPrice)
		if err != nil {
			return entity.SearchResult{}, fmt.Errorf("%s: %w", op, err)
		}
	}

	results.MenuItems = menuResults
	results.Orders = orderResults
	results.Matches = len(menuResults) + len(orderResults)
	return results, nil
}
