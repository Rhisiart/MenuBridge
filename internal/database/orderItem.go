package database

import "fmt"

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

	b := make([]byte, 0, 1+len(m)+len(o)+1+1)

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
	orderItem.quantity = int(data[l-2])
	orderItem.price = int(data[l-1])
}

func (orderItem *OrderItem) Print() {
	fmt.Printf("Order item id: %d\n", orderItem.id)
	fmt.Printf("price: %d\n", orderItem.price)
	fmt.Printf("quantity: %d\n", orderItem.quantity)
}
