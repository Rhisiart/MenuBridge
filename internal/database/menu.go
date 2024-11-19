package database

type Menu struct {
	Id          int        `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	OrderItem   *OrderItem `json:"orderItem,omitempty"`
}

func NewMenu(id int, name string, description string, price float64) Menu {
	return Menu{
		Id:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}
}
