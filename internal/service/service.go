package service

import (
	"context"

	"github.com/Rhisiart/MenuBridge/internal/service/domain"
	"github.com/Rhisiart/MenuBridge/internal/storage"
)

type CategoryService interface {
	FindByOrderId(ctx context.Context, order []byte) ([]byte, error)
}

type FloorService interface {
	FindAll(ctx context.Context) ([]byte, error)
}

type OrderService interface {
	Create(ctx context.Context) ([]byte, error)
	UpsertWithOrderItems(ctx context.Context, data []byte) ([]byte, error)
}

type Service struct {
	CategoryService CategoryService
	FloorService    FloorService
	OrderService    OrderService
}

func NewService(repository *storage.Repository) *Service {
	return &Service{
		CategoryService: domain.NewCategoryService(repository.CategoryRepository),
		FloorService:    domain.NewFloorService(repository.FloorRepository),
		OrderService:    domain.NewOrderService(repository.OrderRepository),
	}
}
