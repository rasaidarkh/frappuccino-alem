package repository

import (
	"context"
	"database/sql"
	"frappuccino-alem/models"
)

type ReportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) *ReportRepository {
	return &ReportRepository{db}
}

func (r *ReportRepository) GetPopularItems(ctx context.Context) ([]models.PopularItem, error) {
	const op = "repository.GetPopularItems"

	// logic here ...

	return nil, nil
}

func (r *ReportRepository) GetTotalSales(ctx context.Context) (float64, error) {
	const op = "repository.GetPopularItems"

	// logic here ...

	return 0, nil
}
