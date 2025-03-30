package service

import (
	"context"
	"fmt"
	"frappuccino-alem/internal/entity"
)

type ReportRepository interface {
	GetPopularItems(ctx context.Context) ([]entity.PopularItem, error)
	GetTotalSales(ctx context.Context) (float64, error)
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
