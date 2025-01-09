package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Rhisiart/MenuBridge/internal/entities"
)

type FloorRepository struct {
	db *sql.DB
}

func NewFloorRepository(db *sql.DB) *FloorRepository {
	return &FloorRepository{
		db: db,
	}
}

func (f *FloorRepository) FindAll(ctx context.Context) ([]*entities.Floor, error) {
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

	rows, err := f.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list := make([]*entities.Floor, 0)

	for rows.Next() {
		var tables []byte
		newFloor := new(entities.Floor)

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
