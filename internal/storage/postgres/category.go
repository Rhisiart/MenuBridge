package postgres

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/Rhisiart/MenuBridge/internal/entities"
)

type CategoryRepository struct {
	db *sql.DB
}

func NewCategoryRepository(db *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		db: db,
	}
}

func (c *CategoryRepository) FindByOrderId(
	ctx context.Context,
	orderId int) ([]*entities.Category, error) {
	query := `SELECT 
		c.id,
		c.name,
		JSON_AGG(
			JSON_BUILD_OBJECT(
				'id', m.id,
				'name', m.name,
				'description', m.description,
				'price', m.price,
				'orderItem', (
					SELECT JSON_BUILD_OBJECT(
						'id', oi.id,
						'quantity', oi.quantity)
					FROM orderitem oi    
					INNER JOIN customerorder o ON o.id = oi.customerorderid 
					WHERE oi.menuid = m.id AND o.id = $1
				)
			)
			ORDER BY m.name
		) AS Menus
	FROM 
		category c
	INNER JOIN 
		menu m ON m.categoryid = c.id
	WHERE 
		c.id IS NOT NULL
	GROUP BY 
		c.id, c.name
	ORDER BY 
		c.id`

	rows, err := c.db.QueryContext(ctx, query, orderId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	list := make([]*entities.Category, 0)

	for rows.Next() {
		var bytes []byte
		newCategory := new(entities.Category)

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
