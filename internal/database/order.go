package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	types "github.com/Rhisiart/MenuBridge/types/interface"
)

type Order struct {
	Id         int         `json:"id,omitempty"`
	CustomerId int         `json:"customerId,omitempty"`
	Amount     float64     `json:"amount,omitempty"`
	Statuscode string      `json:"statuscode,omitempty"`
	CreatedOn  string      `json:"createdOn,omitempty"`
	FloorTable FloorTable  `json:"floorTable,omitempty"`
	OrderItems []OrderItem `json:"orderItems,omitempty"`
}

func NewOrder(id int) *Order {
	return &Order{
		Id: id,
	}
}

func (o *Order) Unmarshal(data []byte) {
	o.Id = int(data[0])
}

func (o *Order) Transaction(ctx context.Context, db *Database) error {
	return db.Transaction(ctx, func(tx *sql.Tx) error {
		var err error

		if o.Id == -1 {
			err = o.Create(ctx, tx)
		} else {
			err = o.Update(ctx, tx)
		}

		if err != nil && err != sql.ErrNoRows {
			tx.Rollback()

			return err
		}

		var transactionValues []string
		var values []interface{}

		for i, orderItem := range o.OrderItems {
			transactionValues = append(transactionValues,
				fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))

			values = append(values, o.Id, orderItem.MenuId, orderItem.Quantity, orderItem.Price)
		}

		query := fmt.Sprintf(`
			INSERT INTO orderitem (customerorderid, menuid, quantity, price)
			VALUES %s
			ON CONFLICT (customerorderid, menuid) 
			DO UPDATE 
			SET quantity = EXCLUDED.quantity, 
				price = EXCLUDED.price`,
			strings.Join(transactionValues, ","))

		_, err = tx.ExecContext(ctx, query, values...)

		return err
	})
}

func (o *Order) Create(ctx context.Context, exec types.Executor) error {
	query := `INSERT INTO customerorder (floortableid, customerid, amount, statuscode, createdon)
				Values($1, $2, $3, $4, $5)
				RETURNING id`

	return exec.QueryRowContext(
		ctx,
		query,
		o.FloorTable.Id,
		o.CustomerId,
		o.Amount,
		o.Statuscode,
		o.CreatedOn).Scan(&o.Id)
}

func (o *Order) Read(ctx context.Context, exec types.Executor) error {
	return nil
}

func (o *Order) ReadAll(ctx context.Context, exec types.Executor) ([]types.Table, error) {
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

	rows, err := exec.QueryContext(ctx, query)

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

func (o *Order) Update(ctx context.Context, exec types.Executor) error {
	query := `UPDATE customerorder
				SET amount = $2
				WHERE id = $1`

	_, err := exec.ExecContext(ctx, query, o.Id, o.Amount)

	return err
}

func (o *Order) Delete(ctx context.Context, exec types.Executor) error {
	return nil
}
