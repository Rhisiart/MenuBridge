package database

import (
	"fmt"
)

const TABLE_SIZE = 2

type Reservation struct {
	Id       int
	Guests   int
	Customer Customer
	Table    Table
	//DateTime string
	//CreateAt string
}

func NewReservation(
	id int,
	customer Customer,
	table Table,
	guests int) Reservation {
	return Reservation{
		Id:       id,
		Customer: customer,
		Table:    table,
		Guests:   guests,
		//DateTime: dateTime,
		//CreateAt: time.Now().String(),
	}
}

func (r *Reservation) MarshalBinary() []byte {
	tableBytes := r.Table.MarshalBinary()
	customerBytes := r.Customer.MarshalBinary()

	//size := 1 + len(customerBytes) + len(tableBytes) + 1 + len(r.DateTime) + len(r.CreateAt)
	size := 1 + 1 + len(customerBytes) + len(tableBytes)

	bytes := make([]byte, 0, size)

	bytes = append(bytes, byte(r.Id))
	bytes = append(bytes, byte(r.Guests))
	bytes = append(bytes, tableBytes...)
	bytes = append(bytes, customerBytes...)
	//bytes = append(bytes, []byte(r.DateTime)...)
	//bytes = append(bytes, []byte(r.CreateAt)...)

	return bytes
}

func (r *Reservation) UnmarshalBinary(data []byte) error {
	var table Table

	err := table.UnmarshalBinary(data[2 : TABLE_SIZE+2])

	if err != nil {
		return fmt.Errorf("unable to unmarshal the table")
	}

	var customer Customer

	err = customer.UnmarshalBinary(data[TABLE_SIZE+2:])

	if err != nil {
		return fmt.Errorf("unable to unmarshal the customer")
	}

	r.Id = int(data[0])
	r.Guests = int(data[1])
	r.Table = table
	r.Customer = customer

	return nil
}

func (r *Reservation) Print() {
	fmt.Printf("-------------------------------------------------\n")
	fmt.Printf("Reservation id: %d\n", r.Id)
	fmt.Printf("Table Id: %d\n", r.Table.id)
	fmt.Printf("Number of guets: %d\n", r.Guests)
	fmt.Printf("Customer Id: %d\n", r.Customer.Id)
	fmt.Printf("Customer Name: %s\n", r.Customer.Name)
}
