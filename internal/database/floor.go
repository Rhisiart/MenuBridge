package database

import (
	"context"
	"database/sql"
	"encoding/json"

	types "github.com/Rhisiart/MenuBridge/types/interface"
)

type Floor struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Tables []Table `json:"tables,omitempty"`
}

func NewFloor(id int, name string) *Floor {
	return &Floor{
		Id:   id,
		Name: name,
	}
}

func (f *Floor) Transaction(ctx context.Context, db *sql.DB) error {
	return nil
}

func (f *Floor) Create(ctx context.Context, exec types.Executor) error {
	return nil
}

func (f *Floor) Read(ctx context.Context, exec types.Executor) error {
	return nil
}

func (f *Floor) ReadAll(ctx context.Context, exec types.Executor) ([]types.Table, error) {
	query := `SELECT 
				f.id,
				f.name,
				JSON_AGG(
					JSON_BUILD_OBJECT(
						'id', ft.id,
						'number', ft.number,
						'capacity', t.capacity,
						'status', ft.Statuscode,
						'order', (
							SELECT JSON_BUILD_OBJECT('id', o.id)
							FROM customerorder o
							INNER JOIN floor_diningtable fdt ON fdt.id = o.floortableid
							WHERE fdt.id = ft.id AND o.statuscode = 'In Progress' AND o.createdon >= CURRENT_DATE
						)
					)
				) AS Tables
				FROM Floor f
				INNER JOIN floor_diningtable ft ON f.id = ft.floorid
				INNER JOIN diningtable t ON t.id = ft.diningtableid
				GROUP BY f.id, f.name
				ORDER BY f.id`

	rows, err := exec.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []types.Table

	for rows.Next() {
		var tables []byte
		newFloor := new(Floor)

		if err := rows.Scan(&newFloor.Id, &newFloor.Name, &tables); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(tables, &newFloor.Tables); err != nil {
			return nil, err
		}

		list = append(list, newFloor)
	}

	return list, nil
}

func (f *Floor) Update(ctx context.Context, exec types.Executor) error {
	return nil
}

func (f *Floor) Delete(ctx context.Context, exec types.Executor) error {
	return nil
}
