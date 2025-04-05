package service

import (
	"context"
	"fmt"
	"time"

	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/dto"
	"frappuccino-alem/internal/store"
)

type InventoryService interface {
	CreateInventoryItem(ctx context.Context, item entity.InventoryItem) (entity.InventoryItem, error)
	GetPaginatedInventoryItems(ctx context.Context, pagination *dto.Pagination) (*dto.PaginationResponse[entity.InventoryItem], error)
	GetPaginatedLeftOverItems(ctx context.Context, pagination *dto.Pagination) (*dto.PaginationResponse[dto.LeftOverItem], error)
	GetInventoryItemById(ctx context.Context, id int64) (entity.InventoryItem, error)
	DeleteInventoryItemById(ctx context.Context, id int64) (entity.InventoryItem, error)
	UpdateInventoryItemById(ctx context.Context, id int64, request dto.InventoryItemRequest) error
}

type inventoryService struct {
	repo store.InventoryRepository
}

func NewInventoryService(repo store.InventoryRepository) InventoryService {
	return &inventoryService{repo: repo}
}

func (s *inventoryService) CreateInventoryItem(ctx context.Context, item entity.InventoryItem) (entity.InventoryItem, error) {
	const op = "service.CreateInventoryItem"

	id, err := s.repo.CreateInventoryItem(ctx, item)
	if err != nil {
		return entity.InventoryItem{}, fmt.Errorf("%s: %w", op, err)
	}

	item.ID = id
	item.CreatedAt = item.CreatedAt.UTC()
	item.UpdatedAt = item.UpdatedAt.UTC()

	return item, nil
}

func (s *inventoryService) GetPaginatedInventoryItems(ctx context.Context, pagination *dto.Pagination) (*dto.PaginationResponse[entity.InventoryItem], error) {
	const op = "service.GetPaginatedInventoryItems"

	totalItems, err := s.repo.GetTotalInventoryCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	totalPages := (totalItems + pagination.PageSize - 1) / pagination.PageSize

	items, err := s.repo.GetAllInventoryItems(ctx, pagination)
	if err != nil {
		if err == store.ErrNotFound {
			return nil, err
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := &dto.PaginationResponse[entity.InventoryItem]{
		CurrentPage: pagination.Page,
		HasNextPage: pagination.Page < totalPages,
		PageSize:    pagination.PageSize,
		TotalPages:  totalPages,
		Data:        items,
	}

	return response, nil
}

func (s *inventoryService) GetInventoryItemById(ctx context.Context, InventoryId int64) (entity.InventoryItem, error) {
	const op = "service.GetInventoryItemById"
	item, err := s.repo.GetInventoryItemById(ctx, InventoryId)
	if err != nil {
		if err == store.ErrNotFound {
			return entity.InventoryItem{}, err
		}
		return entity.InventoryItem{}, fmt.Errorf("%s: %w", op, err)
	}

	return item, nil
}

func (s *inventoryService) DeleteInventoryItemById(ctx context.Context, InventoryId int64) (entity.InventoryItem, error) {
	const op = "service.DeleteInventoryItemById"
	item, err := s.repo.GetInventoryItemById(ctx, InventoryId)
	if err != nil {
		if err == store.ErrNotFound {
			return entity.InventoryItem{}, err
		}
		return entity.InventoryItem{}, fmt.Errorf("%s: %w", op, err)
	}
	_, err = s.repo.DeleteInventoryItemById(ctx, InventoryId)
	if err != nil {
		return entity.InventoryItem{}, fmt.Errorf("%s: %w", op, err)
	}

	return item, nil
}

func (s *inventoryService) UpdateInventoryItemById(ctx context.Context, InventoryId int64, req dto.InventoryItemRequest) error {
	const op = "service.UpdateInventoryItemById"
	return s.repo.UpdateByID(ctx, int64(InventoryId), func(item *entity.InventoryItem) (updated bool, err error) {
		if req.Name != nil {
			if item.ItemName != *req.Name {
				updated = true
				item.ItemName = *req.Name
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

		if req.Price != nil {
			if item.Price != *req.Price {
				updated = true
				item.Price = *req.Price
			}
		}

		if updated {
			item.UpdatedAt = time.Now()
			return
		}

		err = fmt.Errorf("no fields were updated")
		return
	})
}

func (s *inventoryService) GetPaginatedLeftOverItems(ctx context.Context, pagination *dto.Pagination) (*dto.PaginationResponse[dto.LeftOverItem], error) {
	const op = "service.GetPaginatedLeftOverItems"

	totalItems, err := s.repo.GetTotalInventoryCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	totalPages := (totalItems + pagination.PageSize - 1) / pagination.PageSize

	items, err := s.repo.GetAllInventoryItems(ctx, pagination)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	leftoverItems := make([]dto.LeftOverItem, 0)
	for _, item := range items {
		leftoverItems = append(leftoverItems, dto.LeftOverItem{
			Name:     item.ItemName,
			Quantity: item.Quantity,
			UnitType: item.Unit,
			Price:    item.Price,
		})
	}

	response := &dto.PaginationResponse[dto.LeftOverItem]{
		CurrentPage: pagination.Page,
		HasNextPage: pagination.Page < totalPages,
		PageSize:    pagination.PageSize,
		TotalPages:  totalPages,
		Data:        leftoverItems,
	}

	return response, nil
}
