package database

import "fmt"

type Menu struct {
	Id          int
	Name        string
	Description string
	Price       int
	//category
}

func NewMenu(id int, name string, description string, price int) Menu {
	return Menu{
		Id:          id,
		Name:        name,
		Description: description,
		Price:       price,
	}
}

func (m *Menu) MarshalBinary() []byte {
	b := make([]byte, 0, 1+20+len(m.Description)+1)

	name := make([]byte, 20)
	name = append(name, []byte(m.Name)...)

	b = append(b, byte(m.Id))
	b = append(b, name...)
	b = append(b, []byte(m.Description)...)
	b = append(b, byte(m.Price))

	return b
}

func (m *Menu) UnmarshalBinary(data []byte) {
	l := len(data)

	fmt.Printf("data length = %d\n", l)

	m.Id = int(data[0])
	m.Name = string(data[1:21])
	m.Description = string(data[21 : l-1])
	m.Price = int(data[l-1])
}
