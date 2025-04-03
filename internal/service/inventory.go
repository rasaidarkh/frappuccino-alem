package service

import (
	"context"
	"fmt"
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/types"
	"time"
)

type InventoryRepository interface {
	CreateInventoryItem(ctx context.Context, item entity.InventoryItem) (int64, error)
	GetAllInventoryItems(ctx context.Context, pagination *types.Pagination) ([]entity.InventoryItem, error)
	GetTotalInventoryCount(ctx context.Context) (int, error)
	GetInventoryItemById(ctx context.Context, id int64) (entity.InventoryItem, error)
	DeleteInventoryItemById(ctx context.Context, id int64) (int64, error)
	UpdateInventoryItemById(ctx context.Context, id int64, item entity.InventoryItem) (int64, error)
	UpdateByID(ctx context.Context, userID int64, updateFn func(item *entity.InventoryItem) (bool, error)) error
}

type InventoryService struct {
	repo InventoryRepository
}

func NewInventoryService(repo InventoryRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) CreateInventoryItem(ctx context.Context, item entity.InventoryItem) (int64, error) {
	const op = "service.CreateInventoryItem"

	id, err := s.repo.CreateInventoryItem(ctx, item)
	if err != nil {
		return -1, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *InventoryService) GetPaginatedInventoryItems(ctx context.Context, pagination *types.Pagination) (*types.PaginationResponse[entity.InventoryItem], error) {
	totalItems, err := s.repo.GetTotalInventoryCount(ctx)
	if err != nil {
		return nil, err
	}

	totalPages := (totalItems + pagination.PageSize - 1) / pagination.PageSize

	items, err := s.repo.GetAllInventoryItems(ctx, pagination)
	if err != nil {
		return nil, err
	}

	response := &types.PaginationResponse[entity.InventoryItem]{
		CurrentPage: pagination.Page,
		HasNextPage: pagination.Page < totalPages,
		PageSize:    pagination.PageSize,
		TotalPages:  totalPages,
		Data:        items,
	}

	return response, nil
}

func (s *InventoryService) GetInventoryItemById(ctx context.Context, InventoryId int64) (entity.InventoryItem, error) {
	const op = "service.GetInventoryItemById"
	// logic here ...
	item, err := s.repo.GetInventoryItemById(ctx, InventoryId)
	if err != nil {
		return entity.InventoryItem{}, fmt.Errorf("%s: %w", op, err)
	}

	return item, nil
}

func (s *InventoryService) DeleteInventoryItemById(ctx context.Context, InventoryId int64) (entity.InventoryItem, error) {
	const op = "service.DeleteInventoryItemById"
	// logic here ...
	item, err := s.repo.GetInventoryItemById(ctx, InventoryId)
	if err != nil {
		return entity.InventoryItem{}, fmt.Errorf("%s: %w", op, err)
	}
	_, err = s.repo.DeleteInventoryItemById(ctx, InventoryId)
	if err != nil {
		return entity.InventoryItem{}, fmt.Errorf("%s: %w", op, err)
	}

	return item, nil
}

func (s *InventoryService) UpdateInventoryItemById(ctx context.Context, InventoryId int64, req types.InventoryItemRequest) error {
	const op = "service.UpdateInventoryItemById"
	return s.repo.UpdateByID(ctx, int64(InventoryId), func(item *entity.InventoryItem) (updated bool, err error) {
		if req.Name != nil {
			if item.Name != *req.Name {
				updated = true
				item.Name = *req.Name
			}
		}

		if req.Quantity != nil {
			if item.Quantity != *req.Quantity {
				updated = true
				item.Quantity = *req.Quantity
			}
		}

		if req.UnitType != nil {
			if item.Unit != *req.UnitType {
				updated = true
				item.Unit = *req.UnitType
			}
		}

		if updated {
			item.LastUpdated = time.Now()
			return
		}

		err = fmt.Errorf("no fields were updated")
		return
	})
}
