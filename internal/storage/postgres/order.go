package postgres

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Rhisiart/MenuBridge/internal/entities"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (o *OrderRepository) Transaction(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	tx, err := o.db.BeginTx(ctx, nil)

	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			//panic(p)
		}
	}()

	if order.Id == -1 {
		err = o.Create(ctx, order)
	} else {
		err = o.Update(ctx, order)
	}

	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()

		return nil, err
	}

	var transactionValues []string
	var values []interface{}

	for i, orderItem := range order.OrderItems {
		transactionValues = append(transactionValues,
			fmt.Sprintf("($%d, $%d, $%d, $%d)", i*4+1, i*4+2, i*4+3, i*4+4))

		values = append(values, order.Id, orderItem.MenuId, orderItem.Quantity, orderItem.Price)
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

	tx.Commit()

	//TODO: return the result
	return nil, nil
}

func (o *OrderRepository) FindAll(ctx context.Context) ([]*entities.Order, error) {
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

	rows, err := o.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list := make([]*entities.Order, 0)

	for rows.Next() {
		var bytes []byte
		newOrder := new(entities.Order)

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

func (o *OrderRepository) Create(
	ctx context.Context,
	order *entities.Order) error {
	query := `INSERT INTO customerorder (floortableid, customerid, amount, statuscode, createdon)
		Values($1, $2, $3, $4, $5)
		RETURNING id`

	return o.db.QueryRowContext(
		ctx,
		query,
		order.FloorTable.Id,
		order.CustomerId,
		order.Amount,
		order.Statuscode,
		order.CreatedOn).Scan(&order.Id)
}

func (o *OrderRepository) Update(ctx context.Context, order *entities.Order) error {
	query := `UPDATE customerorder
				SET amount = $2
				WHERE id = $1`

	_, err := o.db.ExecContext(ctx, query, order.Id, order.Amount)

	return err
}
