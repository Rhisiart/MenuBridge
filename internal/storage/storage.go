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
	FindAll(ctx context.Context) ([]entities.Floor, error)
}

type Repository struct {
	CategoryRepository CategoryRepository
	FloorRepository    FloorRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		CategoryRepository: postgres.NewCategoryRepository(db),
	}
}
