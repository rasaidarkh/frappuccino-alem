package service

import (
	"context"
	"fmt"
	"frappuccino-alem/models"
)

type MenuRepository interface {
	CreateMenuItem(ctx context.Context, item models.MenuItem) (string, error)
	GetAllMenuItems(ctx context.Context) ([]models.MenuItem, error)
	GetMenuItemById(ctx context.Context, MenuId string) (models.MenuItem, error)
	DeleteMenuItemById(ctx context.Context, id string) error
	UpdateMenuItemById(ctx context.Context, id string, item models.MenuItem) error
}

type MenuService struct {
	repo MenuRepository
}

func NewMenuService(repo MenuRepository) *MenuService {
	return &MenuService{repo: repo}
}

func (s *MenuService) CreateMenuItem(ctx context.Context, item models.MenuItem) (string, error) {
	const op = "service.CreateMenuItem"
	// logic here ...
	id, err := s.repo.CreateMenuItem(ctx, item)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *MenuService) GetAllMenuItems(ctx context.Context) ([]models.MenuItem, error) {
	const op = "service.GetAllMenuItems"
	// logic here ...
	items, err := s.repo.GetAllMenuItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return items, nil
}

func (s *MenuService) GetMenuItemById(ctx context.Context, id string) (models.MenuItem, error) {
	const op = "service.GetMenuItemById"
	// logic here ...
	item, err := s.repo.GetMenuItemById(ctx, id)
	if err != nil {
		return models.MenuItem{}, fmt.Errorf("%s: %w", op, err)
	}

	return item, nil
}

func (s *MenuService) DeleteMenuItemById(ctx context.Context, id string) error {
	const op = "service.DeleteMenuItemById"
	// logic here ...
	err := s.repo.DeleteMenuItemById(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *MenuService) UpdateMenuItemById(ctx context.Context, id string, item models.MenuItem) error {
	const op = "service.UpdateMenuItemById"
	// logic here ...
	err := s.repo.UpdateMenuItemById(ctx, id, item)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
