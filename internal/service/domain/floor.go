package domain

import (
	"context"

	"github.com/Rhisiart/MenuBridge/internal/entities"
	"github.com/Rhisiart/MenuBridge/internal/storage"
)

type FloorService struct {
	floorRepository storage.FloorRepository
}

func NewFloorService(floorRepo storage.FloorRepository) *FloorService {
	return &FloorService{
		floorRepository: floorRepo,
	}
}

func (f *FloorService) FindAll(ctx context.Context) ([]*entities.Floor, error) {
	return nil, nil
}
