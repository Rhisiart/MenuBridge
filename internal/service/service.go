package service

import (
	"context"

	"github.com/Rhisiart/MenuBridge/internal/entities"
	"github.com/Rhisiart/MenuBridge/internal/service/domain"
	"github.com/Rhisiart/MenuBridge/internal/storage"
)

type CategoryService interface {
	FindByOrderId(ctx context.Context, orderId int) ([]*entities.Category, error)
}

type FloorService interface {
	FindAll(ctx context.Context) ([]*entities.Floor, error)
}

type Service struct {
	CategoryService CategoryService
	FloorService    FloorService
}

func NewService(repository *storage.Repository) *Service {
	return &Service{
		CategoryService: domain.NewCategoryService(repository.CategoryRepository),
		FloorService:    domain.NewFloorService(repository.FloorRepository),
	}
}
