package store

import (
	"context"
	"database/sql"
	"frappuccino-alem/models"
)

type ReportStore struct {
	db *sql.DB
}

func NewReportStore(db *sql.DB) *ReportStore {
	return &ReportStore{db}
}

func (r *ReportStore) GetPopularItems(ctx context.Context) ([]models.PopularItem, error) {
	const op = "Store.GetPopularItems"

	// logic here ...

	return nil, nil
}

func (r *ReportStore) GetTotalSales(ctx context.Context) (float64, error) {
	const op = "Store.GetPopularItems"

	// logic here ...

	return 0, nil
}
