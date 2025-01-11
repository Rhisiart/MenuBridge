package domain

import (
	"context"
	"encoding/json"

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
	order []byte) ([]byte, error) {
	newOrder := new(entities.Order)
	newOrder.Unmarshal(order)

	ctgr, err := c.categoryRepository.FindByOrderId(ctx, newOrder.Id)

	if err != nil {
		return nil, err
	}

	return json.Marshal(ctgr)
}
