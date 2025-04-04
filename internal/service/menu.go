package service

import (
	"context"
	"fmt"
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/types"
)

type MenuRepository interface {
	CreateMenuItem(ctx context.Context, item entity.MenuItem) (string, error)
	GetAllMenuItems(ctx context.Context, pagination *types.Pagination) ([]entity.MenuItem, error)
	GetTotalMenuCount(ctx context.Context) (int, error)
	GetMenuItemById(ctx context.Context, MenuId string) (entity.MenuItem, error)
	DeleteMenuItemById(ctx context.Context, id string) error
	UpdateMenuItemById(ctx context.Context, id string, item entity.MenuItem) error
}

type MenuService struct {
	repo MenuRepository
}

func NewMenuService(repo MenuRepository) *MenuService {
	return &MenuService{repo: repo}
}

func (s *MenuService) CreateMenuItem(ctx context.Context, item entity.MenuItem) (string, error) {
	const op = "service.CreateMenuItem"
	id, err := s.repo.CreateMenuItem(ctx, item)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *MenuService) GetPaginatedMenuItems(ctx context.Context, pagination *types.Pagination) (*types.PaginationResponse[entity.MenuItem], error) {
	const op = "service.GetPaginatedMenuItems"

	totalItems, err := s.repo.GetTotalMenuCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	totalPages := (totalItems + pagination.PageSize - 1) / pagination.PageSize

	items, err := s.repo.GetAllMenuItems(ctx, pagination)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	response := &types.PaginationResponse[entity.MenuItem]{
		CurrentPage: pagination.Page,
		HasNextPage: pagination.Page < totalPages,
		PageSize:    pagination.PageSize,
		TotalPages:  totalPages,
		Data:        items,
	}

	return response, nil
}

func (s *MenuService) GetMenuItemById(ctx context.Context, id string) (entity.MenuItem, error) {
	const op = "service.GetMenuItemById"
	// logic here ...
	item, err := s.repo.GetMenuItemById(ctx, id)
	if err != nil {
		return entity.MenuItem{}, fmt.Errorf("%s: %w", op, err)
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

func (s *MenuService) UpdateMenuItemById(ctx context.Context, id string, item entity.MenuItem) error {
	const op = "service.UpdateMenuItemById"
	// logic here ...
	err := s.repo.UpdateMenuItemById(ctx, id, item)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
