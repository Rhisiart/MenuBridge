package database

import (
	"context"
	"database/sql"
	"encoding/json"

	types "github.com/Rhisiart/MenuBridge/types/interface"
)

type Order struct {
	Id         int         `json:"id,omitempty"`
	CustomerId int         `json:"customerId,omitempty"`
	Amount     float64     `json:"amount,omitempty"`
	Statuscode string      `json:"statuscode,omitempty"`
	CreatedOn  string      `json:"createdOn,omitempty"`
	FloorTable FloorTable  `json:"floorTable,omitempty"`
	OrderItem  []OrderItem `json:"orderItems,omitempty"`
}

func NewOrder(id int) *Order {
	return &Order{
		Id: id,
	}
}

func (o *Order) Unmarshal(data []byte) {
	o.Id = int(data[0])
}

func (o *Order) Transaction(ctx context.Context, db *sql.DB) error {
	var err error

	if o.Id == -1 {
		err = o.Create(ctx, db)
	} else {
		err = o.Update(ctx, db)
	}

	return err
}

func (o *Order) Create(ctx context.Context, db *sql.DB) error {
	query := `WITH floor_table AS (
					SELECT ft.id
					FROM floor_diningtable ft
					WHERE ft.floorid = $1 AND ft.diningtableid = $2 AND ft.number = $3
				)
				INSERT INTO customerorder (floortableid, customerid, amount, statuscode, createdon)
				SELECT id, $4, $5, $6, $7
				FROM floor_table
				RETURNING id`

	err := db.QueryRowContext(
		ctx,
		query,
		o.FloorTable.FloorId,
		o.FloorTable.TableId,
		o.FloorTable.Number,
		o.CustomerId,
		o.Amount,
		o.Statuscode,
		o.CreatedOn).Scan(&o.Id)

	return err
}

func (o *Order) Read(ctx context.Context, db *sql.DB) error {
	return nil
}

func (o *Order) ReadAll(ctx context.Context, db *sql.DB) ([]types.Table, error) {
	query := `SELECT 
				o.id, 
				o.amount, 
				o.statuscode,
				o.customerid,
				o.createdon,
				JSON_BUILD_OBJECT(
					'id', ft.id,
					'number', ft.number,
					'tableId', ft.diningtableid,
					'floorId', ft.floorid
				) AS floorTable
			FROM
				customerorder o
			INNER JOIN
				floor_diningtable ft ON ft.id = o.floortableid
			WHERE
				o.createdon >= CURRENT_DATE AND o.statuscode = 'In Progress'
			GROUP BY
				o.id, ft.id`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []types.Table

	for rows.Next() {
		var bytes []byte
		newOrder := new(Order)

		if err := rows.Scan(
			&newOrder.Id,
			&newOrder.Amount,
			&newOrder.Statuscode,
			&newOrder.CustomerId,
			&newOrder.CreatedOn,
			&bytes); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(bytes, &newOrder.FloorTable); err != nil {
			return nil, err
		}

		list = append(list, newOrder)
	}

	return list, nil
}

func (o *Order) Update(ctx context.Context, db *sql.DB) error {
	return nil
}

func (o *Order) Delete(ctx context.Context, db *sql.DB) error {
	return nil
}
