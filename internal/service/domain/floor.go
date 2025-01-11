package domain

import (
	"context"
	"encoding/json"

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

func (f *FloorService) FindAll(ctx context.Context) ([]byte, error) {
	floors, err := f.floorRepository.FindAll(ctx)

	if err != nil {
		return nil, err
	}

	return json.Marshal(floors)
}
