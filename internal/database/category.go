package database

import (
	"context"
	"database/sql"
	"encoding/json"

	types "github.com/Rhisiart/MenuBridge/types/interface"
)

type Category struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Menus []Menu `json:"menus"`
}

func NewCategory(id int, name string) *Category {
	return &Category{
		Id:   id,
		Name: name,
	}
}

func (c Category) Create(ctx context.Context, db *sql.DB) error {
	return nil
}

func (c Category) Read(ctx context.Context, db *sql.DB) error {
	return nil
}

func (c Category) ReadAll(ctx context.Context, db *sql.DB) ([]types.Table, error) {
	query := `SELECT 
				c.id,
				c.name,
				JSON_AGG(
					JSON_BUILD_OBJECT(
						'id', m.id,
						'name', m.name,
						'description', m.description,
						'price', m.price
					)
					ORDER BY m.name
				) AS Menus
				FROM 
					category c
				INNER JOIN 
					menu m  ON c.id = m.categoryid
				GROUP BY 
					c.id, c.name
				ORDER BY
					c.id;`

	rows, err := db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var list []types.Table

	for rows.Next() {
		var bytes []byte
		newCategory := new(Category)

		if err := rows.Scan(&newCategory.Id, &newCategory.Name, &bytes); err != nil {
			return nil, err
		}

		if err := json.Unmarshal(bytes, &newCategory.Menus); err != nil {
			return nil, err
		}

		list = append(list, newCategory)
	}

	return list, nil
}

func (c Category) Update(ctx context.Context, db *sql.DB) error {
	return nil
}

func (c Category) Delete(ctx context.Context, db *sql.DB) error {
	return nil
}
