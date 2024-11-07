package database

type Order struct {
	Id       int
	persons  int
	table    Table
	customer Customer
	//Status       enum.OrderStatus
}

func NewOrder(id int, persons int, table Table, customer Customer) Order {
	return Order{
		Id:       id,
		persons:  persons,
		table:    table,
		customer: customer,
		//Status:       enum.Preparing,
	}
}
