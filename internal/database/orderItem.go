package database

type OrderItem struct {
	id       int
	order    Order
	menu     Menu
	quantity int
	price    int
}

func NewOrderItem(id int, order Order, menu Menu, quantity int, price int) OrderItem {
	return OrderItem{
		id:       id,
		order:    order,
		menu:     menu,
		quantity: quantity,
		price:    price,
	}
}
