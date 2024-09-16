package database

import "fmt"

type Order struct {
	Persons  int
	Table    Table
	Customer Customer
	//Status       enum.OrderStatus
}

func NewOrder(persons int, table Table, customer Customer) *Order {
	return &Order{
		Persons:  persons,
		Table:    table,
		Customer: customer,
		//Status:       enum.Preparing,
	}
}

func (o *Order) MarshalBinary() []byte {
	q := o.Table.MarshalBinary()
	c := o.Customer.MarshalBinary()

	bytes := make([]byte, 0, 1+len(q)+len(c))

	bytes = append(bytes, byte(o.Persons))
	bytes = append(bytes, q...)
	bytes = append(bytes, c...)

	return bytes
}

func (o *Order) UnmarshalBinary(data []byte) error {
	var t Table

	err := t.UnmarshalBinary(data[1:3])

	if err != nil {
		return fmt.Errorf("error when unmarshal table in order")
	}

	var c Customer

	err = c.UnmarshalBinary(data[3:])

	if err != nil {
		return fmt.Errorf("error when unmarshal reservation in order")
	}

	o.Customer = c
	o.Table = t
	o.Persons = int(data[0])

	return nil
}
