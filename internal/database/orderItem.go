package database

type OrderItem struct {
	id       int
	menu     Menu
	order    Order
	quantity int
	price    int
}

func NewOrderItem(id int, menu Menu, order Order, quantity int) OrderItem {
	return OrderItem{
		id:       id,
		menu:     menu,
		order:    order,
		quantity: quantity,
		price:    quantity * menu.Price,
	}
}

func (orderItem *OrderItem) MarshalBinary() []byte {
	m := orderItem.menu.MarshalBinary()
	o := orderItem.order.MarshalBinary()

	b := make([]byte, 1+len(m)+len(o)+1+1)

	b = append(b, byte(orderItem.id))
	b = append(b, m...)
	b = append(b, o...)
	b = append(b, byte(orderItem.quantity))
	b = append(b, byte(orderItem.price))

	return b
}

func (orderItem *OrderItem) UnmarshalBinary(data []byte) {
	l := len(data)

	var menu Menu

	menu.UnmarshalBinary(data[1 : 72+1])

	var order Order

	order.UnmarshalBinary(data[73 : l-3])

	orderItem.id = int(data[0])
	orderItem.menu = menu
	orderItem.order = order
	orderItem.price = int(data[l-2])
	orderItem.quantity = int(data[l-1])
}
