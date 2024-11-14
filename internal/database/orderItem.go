package database

type OrderItem struct {
	id       int
	menu     Menu
	order    Order
	quantity int
	price    float32
}

func NewOrderItem(id int, menu Menu, order Order, quantity int) OrderItem {
	return OrderItem{
		id:       id,
		menu:     menu,
		order:    order,
		quantity: quantity,
		price:    float32(quantity) * menu.Price,
	}
}
