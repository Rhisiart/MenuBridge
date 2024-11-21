package database

import (
	"context"
	"database/sql"

	types "github.com/Rhisiart/MenuBridge/types/interface"
)

type OrderItem struct {
	Id       int     `json:"id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

func (oi OrderItem) Create(ctx context.Context, db *sql.DB) error {
	return nil
}

func (oi OrderItem) Read(ctx context.Context, db *sql.DB) error {
	return nil
}

func (oi OrderItem) ReadAll(ctx context.Context, db *sql.DB) ([]types.Table, error) {
	return nil, nil
}

func (oi OrderItem) Update(ctx context.Context, db *sql.DB) error {
	query := `UPDATE orderitem
				SET quantity = $1, price = $2
				WHERE id = $3`

	_, err := db.ExecContext(ctx, query, oi.Quantity, oi.Price, oi.Id)

	return err
}

func (oi OrderItem) Delete(ctx context.Context, db *sql.DB) error {
	return nil
}
