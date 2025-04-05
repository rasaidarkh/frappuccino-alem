package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"frappuccino-alem/internal/entity"
	"frappuccino-alem/internal/handlers/dto"
	"frappuccino-alem/internal/store"
)

type MenuService interface {
	CreateMenuItem(ctx context.Context, item entity.MenuItem) (entity.MenuItem, error)
	GetPaginatedMenuItems(ctx context.Context, pagination *dto.Pagination) (*dto.PaginationResponse[entity.MenuItem], error)
	GetMenuItemById(ctx context.Context, id int64) (entity.MenuItem, error)
	DeleteMenuItemById(ctx context.Context, id int64) error
	UpdateMenuItemById(ctx context.Context, id int64, request dto.MenuItemRequest) error
}

type menuService struct {
	menuRepo      store.MenuRepository
	inventoryRepo store.InventoryRepository
}

func NewMenuService(
	menuRepo store.MenuRepository,
	inventoryRepo store.InventoryRepository,
) MenuService {
	return &menuService{
		menuRepo:      menuRepo,
		inventoryRepo: inventoryRepo,
	}
}

func (s *menuService) CreateMenuItem(ctx context.Context, item entity.MenuItem) (entity.MenuItem, error) {
	const op = "service.CreateMenuItem"

	// Validate ingredients exist
	for _, ing := range item.Ingredients {
		_, err := s.inventoryRepo.GetInventoryItemById(ctx, ing.ID)
		if err != nil {
			if errors.Is(err, store.ErrNotFound) {
				return entity.MenuItem{}, fmt.Errorf("ingredient %d not found", ing.ID)
			}
			return entity.MenuItem{}, fmt.Errorf("%s: %w", op, err)
		}
	}

	// Proceed with creation
	id, err := s.menuRepo.CreateMenuItem(ctx, item)
	if err != nil {
		return item, fmt.Errorf("%s: %w", op, err)
	}

	item.ID = id
	return item, nil
}

func (s *menuService) GetPaginatedMenuItems(ctx context.Context, pagination *dto.Pagination) (*dto.PaginationResponse[entity.MenuItem], error) {
	const op = "service.GetPaginatedMenuItems"

	totalItems, err := s.menuRepo.GetTotalMenuCount(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	totalPages := (totalItems + pagination.PageSize - 1) / pagination.PageSize

	items, err := s.menuRepo.GetAllMenuItems(ctx, pagination)
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

func (s *menuService) GetMenuItemById(ctx context.Context, id int64) (entity.MenuItem, error) {
	const op = "service.GetMenuItemById"

	item, err := s.menuRepo.GetMenuItemById(ctx, id)
	if err != nil {
		return entity.MenuItem{}, fmt.Errorf("%s: %w", op, err)
	}
	return item, nil
}

func (s *menuService) UpdateMenuItemById(ctx context.Context, id int64, req dto.MenuItemRequest) error {
	const op = "service.UpdateMenuItemById"

	return s.menuRepo.UpdateByID(ctx, int64(id), func(item *entity.MenuItem) (updated bool, err error) {
		if req.Name != nil {
			if item.Name != *req.Name {
				updated = true
				item.Name = *req.Name
			}
		}

		if req.Price != nil {
			if item.Price != *req.Price {
				updated = true
				item.Price = *req.Price
			}
		}

		if req.Description != nil {
			if item.Description != *req.Description {
				updated = true
				item.Description = *req.Description
			}
		}

		if req.Price != nil {
			if item.Price != *req.Price {
				updated = true
				item.Price = *req.Price
			}
		}

		if req.Categories != nil {
			if !testEq(item.Categories, *req.Categories) {
				updated = true
				item.Categories = *req.Categories
			}
		}

		if req.Allergens != nil {
			if !testEq(item.Allergens, *req.Allergens) {
				updated = true
				item.Allergens = *req.Allergens
			}
		}

		if req.Metadata != nil {
			updated = true
			item.Metadata = *req.Metadata
		}

		if req.Ingredients != nil {
			updated = true
			item.Ingredients = req.MapToEntity().Ingredients
		}

		if updated {
			item.UpdatedAt = time.Now()
			return
		}

		err = fmt.Errorf("no fields were updated")
		return
	})
}

func (s *menuService) DeleteMenuItemById(ctx context.Context, id int64) error {
	const op = "service.DeleteMenuItemById"

	if err := s.menuRepo.DeleteMenuItemById(ctx, id); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}

func testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
