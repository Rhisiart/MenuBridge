package database

type Menu struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	//category
}

func NewMenu(id int, name string, description string, price float32) Menu {
	return Menu{
		Id:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}
}
