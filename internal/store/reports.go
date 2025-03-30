package store

import (
	"context"
	"database/sql"
	"frappuccino-alem/internal/entity"
)

type ReportStore struct {
	db *sql.DB
}

func NewReportStore(db *sql.DB) *ReportStore {
	return &ReportStore{db}
}

func (r *ReportStore) GetPopularItems(ctx context.Context) ([]entity.PopularItem, error) {
	const op = "Store.GetPopularItems"

	// logic here ...

	return nil, nil
}

func (r *ReportStore) GetTotalSales(ctx context.Context) (float64, error) {
	const op = "Store.GetPopularItems"

	// logic here ...

	return 0, nil
}
