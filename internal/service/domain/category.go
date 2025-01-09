package domain

import (
	"context"

	"github.com/Rhisiart/MenuBridge/internal/entities"
	"github.com/Rhisiart/MenuBridge/internal/storage"
)

type CategoryService struct {
	categoryRepository storage.CategoryRepository
}

func NewCategoryService(categoryRepo storage.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepository: categoryRepo,
	}
}

func (c *CategoryService) FindByOrderId(
	ctx context.Context,
	orderId int) ([]*entities.Category, error) {
	return nil, nil
}
