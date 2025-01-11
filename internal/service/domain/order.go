package domain

import (
	"context"
	"encoding/json"

	"github.com/Rhisiart/MenuBridge/internal/entities"
	"github.com/Rhisiart/MenuBridge/internal/storage"
)

type OrderService struct {
	OrderRepository storage.OrderRepository
}

func NewOrderService(orderRepo storage.OrderRepository) *OrderService {
	return &OrderService{
		OrderRepository: orderRepo,
	}
}

func (o *OrderService) Create(ctx context.Context) ([]byte, error) {
	order := new(entities.Order)
	o.OrderRepository.Create(ctx, order)

	return json.Marshal(order)
}

func (o *OrderService) UpsertWithOrderItems(ctx context.Context, data []byte) ([]byte, error) {
	order := new(entities.Order)
	err := json.Unmarshal(data, order)

	if err != nil {
		return nil, err
	}

	newOrder, err := o.OrderRepository.UpsertOrderWithOrderItems(ctx, order)

	if err != nil {
		return nil, err
	}

	return json.Marshal(newOrder)
}

func (o *OrderService) UpdateStatus(ctx context.Context, data []byte) error {
	order := new(entities.Order)
	err := json.Unmarshal(data, order)

	if err != nil {
		return err
	}

	return o.OrderRepository.UpdateStatus(ctx, order)
}
