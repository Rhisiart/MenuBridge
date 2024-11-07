package database

import (
	"context"
	"database/sql"

	types "github.com/Rhisiart/MenuBridge/types/interface"
)

type Floor struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Tables []uint8 `json:"tables"`
}

func NewFloor(id int, name string) Floor {
	return Floor{
		Id:   id,
		Name: name,
	}
}

func (f Floor) Create(ctx context.Context, db *sql.DB) error {
	return nil
}

func (f Floor) Read(ctx context.Context, db *sql.DB) error {
	return nil
}

func (f Floor) ReadAll(ctx context.Context, db *sql.DB, list []types.Table) error {
	query := `SELECT 
				f.id,
				f.name,
				JSON_AGG(
					JSON_BUILD_OBJECT(
						'id', dt.id,
						'number', t.number,
						'capacity', dt.capacity,
						'status', t.Statuscode
					)
				) AS Tables
			FROM 
				Floor f
			LEFT JOIN 
				floor_diningtable t ON f.id = t.floorid
			LEFT JOIN
				diningtable dt ON dt.id = t.diningtableid
			GROUP BY 
				f.id, f.name
			ORDER BY
				f.id;`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		newFloor := new(Floor)

		if err := rows.Scan(&newFloor.Id, &newFloor.Name, &newFloor.Tables); err != nil {
			return err
		}

		list = append(list, newFloor)
	}

	return nil
}

func (f Floor) Update(ctx context.Context, db *sql.DB) error {
	return nil
}

func (f Floor) Delete(ctx context.Context, db *sql.DB) error {
	return nil
}
