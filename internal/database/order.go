package database

import "fmt"

type Order struct {
	id       int
	persons  int
	table    Table
	customer Customer
	//Status       enum.OrderStatus
}

func NewOrder(id int, persons int, table Table, customer Customer) Order {
	return Order{
		id:       id,
		persons:  persons,
		table:    table,
		customer: customer,
		//Status:       enum.Preparing,
	}
}

func (o *Order) MarshalBinary() []byte {
	q := o.table.MarshalBinary()
	c := o.customer.MarshalBinary()

	bytes := make([]byte, 0, 1+1+len(q)+len(c))

	bytes = append(bytes, byte(o.id))
	bytes = append(bytes, byte(o.persons))
	bytes = append(bytes, q...)
	bytes = append(bytes, c...)

	return bytes
}

func (o *Order) UnmarshalBinary(data []byte) error {
	var t Table

	err := t.UnmarshalBinary(data[2:4])

	if err != nil {
		return fmt.Errorf("error when unmarshal table in order")
	}

	var c Customer

	err = c.UnmarshalBinary(data[4:])

	if err != nil {
		return fmt.Errorf("error when unmarshal reservation in order")
	}

	o.id = int(data[0])
	o.persons = int(data[1])
	o.customer = c
	o.table = t

	return nil
}

func (o *Order) Print() {
	fmt.Printf("-------------------------------------------------\n")
	fmt.Printf("Order id: %d\n", o.id)
	fmt.Printf("Table Id: %d\n", o.table.id)
	fmt.Printf("Number of persons: %d\n", o.persons)
	fmt.Printf("Customer Id: %d\n", o.customer.Id)
	fmt.Printf("Customer Name: %s\n", o.customer.Name)
}
