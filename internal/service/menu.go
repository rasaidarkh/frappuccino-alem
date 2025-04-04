package service

import (
	"context"
	"fmt"
	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/dto"
)

type MenuRepository interface {
	CreateMenuItem(ctx context.Context, item entity.MenuItem) (int64, error)
	GetAllMenuItems(ctx context.Context, pagination *dto.Pagination) ([]entity.MenuItem, error)
	GetTotalMenuCount(ctx context.Context) (int, error)
	GetMenuItemById(ctx context.Context, menuID int64) (entity.MenuItem, error)
	DeleteMenuItemById(ctx context.Context, id int64) error
	UpdateMenuItemById(ctx context.Context, id int64, item entity.MenuItem) error
}

type MenuService struct {
	repo MenuRepository
}

func NewMenuService(repo MenuRepository) *MenuService {
	return &MenuService{repo: repo}
}

func (s *MenuService) CreateMenuItem(ctx context.Context, item entity.MenuItem) (int64, error) {
	const op = "service.CreateMenuItem"

	id, err := s.repo.CreateMenuItem(ctx, item)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	return id, nil
}

func (s *MenuService) GetPaginatedMenuItems(ctx context.Context, pagination *dto.Pagination) (*dto.PaginationResponse[entity.MenuItem], error) {
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

	return &dto.PaginationResponse[entity.MenuItem]{
		CurrentPage: pagination.Page,
		HasNextPage: pagination.Page < totalPages,
		PageSize:    pagination.PageSize,
		TotalPages:  totalPages,
		Data:        items,
	}, nil
}

func (s *MenuService) GetMenuItemById(ctx context.Context, id int64) (entity.MenuItem, error) {
	const op = "service.GetMenuItemById"

	item, err := s.repo.GetMenuItemById(ctx, id)
	if err != nil {
		return entity.MenuItem{}, fmt.Errorf("%s: %w", op, err)
	}
	return item, nil
}

func (s *MenuService) DeleteMenuItemById(ctx context.Context, id int64) error {
	const op = "service.DeleteMenuItemById"

	if err := s.repo.DeleteMenuItemById(ctx, id); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func (s *MenuService) UpdateMenuItemById(ctx context.Context, id int64, item entity.MenuItem) error {
	const op = "service.UpdateMenuItemById"

	if err := s.repo.UpdateMenuItemById(ctx, id, item); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
