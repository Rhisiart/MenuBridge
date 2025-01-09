package database

type Menu struct {
	Id          int        `json:"id" db:"id"`
	Name        string     `json:"name,omitempty" db:"name"`
	Description string     `json:"description,omitempty" db:"description"`
	Price       float64    `json:"price,omitempty" db:"price"`
	OrderItem   *OrderItem `json:"orderItem,omitempty"`
}

func NewMenu(id int, name string) Menu {
	return Menu{
		Id:   id,
		Name: name,
	}
}
