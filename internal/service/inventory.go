package service

import (
	"context"
	"fmt"
	"frappuccino-alem/internal/entity"
)

type InventoryRepository interface {
	GetAllInventoryItems(ctx context.Context) ([]entity.InventoryItem, error)
	GetInventoryItemById(ctx context.Context, id int64) (entity.InventoryItem, error)
	DeleteInventoryItemById(ctx context.Context, id int64) (int64, error)
	UpdateInventoryItemById(ctx context.Context, id int64, item entity.InventoryItem) (int64, error)
	CreateInventoryItem(ctx context.Context, item entity.InventoryItem) (int64, error)
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

func (s *InventoryService) GetAllInventoryItems(ctx context.Context) ([]entity.InventoryItem, error) {
	const op = "service.GetAllInventoryItems"
	// logic here ...
	items, err := s.repo.GetAllInventoryItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return items, nil
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

func (s *InventoryService) DeleteInventoryItemById(ctx context.Context, InventoryId int64) error {
	const op = "service.DeleteInventoryItemById"
	// logic here ...
	_, err := s.repo.DeleteInventoryItemById(ctx, InventoryId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *InventoryService) UpdateInventoryItemById(ctx context.Context, InventoryId int64, item entity.InventoryItem) error {
	const op = "service.UpdateInventoryItemById"
	// logic here ...
	_, err := s.repo.UpdateInventoryItemById(ctx, InventoryId, item)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
