package storage

import (
	"context"
	"database/sql"

	"github.com/Rhisiart/MenuBridge/internal/entities"
	"github.com/Rhisiart/MenuBridge/internal/storage/postgres"
)

type MenuRepository interface {
}

type CategoryRepository interface {
	FindByOrderId(ctx context.Context, orderId int) ([]*entities.Category, error)
}

type FloorRepository interface {
	FindAll(ctx context.Context) ([]*entities.Floor, error)
}

type OrderRepository interface {
	FindAll(ctx context.Context) ([]*entities.Order, error)
	Create(ctx context.Context, order *entities.Order) error
	UpdateAmount(ctx context.Context, order *entities.Order) error
	UpdateStatus(ctx context.Context, order *entities.Order) error
	UpsertOrderWithOrderItems(ctx context.Context, order *entities.Order) (*entities.Order, error)
}

type Repository struct {
	CategoryRepository CategoryRepository
	FloorRepository    FloorRepository
	OrderRepository    OrderRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		CategoryRepository: postgres.NewCategoryRepository(db),
		FloorRepository:    postgres.NewFloorRepository(db),
		OrderRepository:    postgres.NewOrderRepository(db),
	}
}
