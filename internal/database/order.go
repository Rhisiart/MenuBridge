package database

import (
	"context"
	"database/sql"

	types "github.com/Rhisiart/MenuBridge/types/interface"
)

type Order struct {
	Id       int `json:"id"`
	customer Customer
}

func NewOrder(id int, customer Customer) Order {
	return Order{
		Id:       id,
		customer: customer,
	}
}

func (o *Order) Unmarshal(data []byte) {
	o.Id = int(data[0])
}

func (o *Order) Create(ctx context.Context, db *sql.DB) error {
	return nil
}

func (o *Order) Read(ctx context.Context, db *sql.DB) error {
	return nil
}

func (o *Order) ReadAll(ctx context.Context, db *sql.DB) ([]types.Table, error) {
	return nil, nil
}

func (o *Order) Update(ctx context.Context, db *sql.DB) error {
	return nil
}

func (o *Order) Delete(ctx context.Context, db *sql.DB) error {
	return nil
}
