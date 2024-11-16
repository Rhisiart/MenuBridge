package database

import (
	"context"
	"database/sql"

	types "github.com/Rhisiart/MenuBridge/types/interface"
)

type Order struct {
	Id         int `json:"id"`
	tableId    int
	floorId    int
	Amount     float64 `json:"amount"`
	Statuscode string  `json:"statuscode"`
	customer   Customer
}

func NewOrder(id int, customer Customer) Order {
	return Order{
		Id:       id,
		customer: customer,
	}
}

func (o *Order) Unmarshal(data []byte) {
	o.floorId = int(data[0])
	o.tableId = int(data[1])
}

func (o *Order) Create(ctx context.Context, db *sql.DB) error {
	return nil
}

func (o *Order) Read(ctx context.Context, db *sql.DB) error {
	query := `SELECT o.id, o.amount, o.statuscode
				FROM
					customerorder o
				INNER JOIN
					floor_diningtable ft 
				ON 
					ft.id = o.floortableid
				INNER JOIN 
					floor f
				ON 
					f.id = ft.floorid
				INNER JOIN
					diningtable t
				ON
					t.id = ft.diningtableid        
				WHERE
					t.id = $1 AND f.id = $2`

	err := db.QueryRowContext(
		ctx,
		query,
		o.tableId,
		o.floorId).Scan(
		&o.Id,
		&o.Amount,
		&o.Statuscode)

	if err != nil {
		return err
	}

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
