package service

import (
	"context"
	"fmt"
	"frappuccino-alem/models"
)

type InventoryRepository interface {
	GetAllInventoryItems(ctx context.Context) ([]models.InventoryItem, error)
	GetInventoryItemById(ctx context.Context, id string) (models.InventoryItem, error)
	DeleteInventoryItemById(ctx context.Context, id string) error
	UpdateInventoryItemById(ctx context.Context, id string, item models.InventoryItem) error
	CreateInventoryItem(ctx context.Context, item models.InventoryItem) (string, error)
}

type InventoryService struct {
	repo InventoryRepository
}

func NewInventoryService(repo InventoryRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (s *InventoryService) CreateInventoryItem(ctx context.Context, item models.InventoryItem) (string, error) {
	const op = "service.CreateInventoryItem"
	// logic here ...
	id, err := s.repo.CreateInventoryItem(ctx, item)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *InventoryService) GetAllInventoryItems(ctx context.Context) ([]models.InventoryItem, error) {
	const op = "service.GetAllInventoryItems"
	// logic here ...
	items, err := s.repo.GetAllInventoryItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return items, nil
}

func (s *InventoryService) GetInventoryItemById(ctx context.Context, InventoryId string) (models.InventoryItem, error) {
	const op = "service.GetInventoryItemById"
	// logic here ...
	item, err := s.repo.GetInventoryItemById(ctx, InventoryId)
	if err != nil {
		return models.InventoryItem{}, fmt.Errorf("%s: %w", op, err)
	}

	return item, nil
}

func (s *InventoryService) DeleteInventoryItemById(ctx context.Context, InventoryId string) error {
	const op = "service.DeleteInventoryItemById"
	// logic here ...
	err := s.repo.DeleteInventoryItemById(ctx, InventoryId)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *InventoryService) UpdateInventoryItemById(ctx context.Context, InventoryId string, item models.InventoryItem) error {
	const op = "service.UpdateInventoryItemById"
	// logic here ...
	err := s.repo.UpdateInventoryItemById(ctx, InventoryId, item)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
