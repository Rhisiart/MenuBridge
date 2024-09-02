package database

import "fmt"

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
	customerBytes := r.Customer.MarshalBinary()
	tableBytes := r.Table.MarshalBinary()

	//size := 1 + len(customerBytes) + len(tableBytes) + 1 + len(r.DateTime) + len(r.CreateAt)
	size := 1 + 1 + len(customerBytes) + len(tableBytes)

	bytes := make([]byte, 0, size)

	bytes = append(bytes, byte(r.Id))
	bytes = append(bytes, byte(r.Guests))
	bytes = append(bytes, customerBytes...)
	bytes = append(bytes, tableBytes...)
	//bytes = append(bytes, []byte(r.DateTime)...)
	//bytes = append(bytes, []byte(r.CreateAt)...)

	return bytes
}

func (r *Reservation) UnmarshalBinary(data []byte) error {
	var table Table

	err := table.UnmarshalBinary(data[0:TABLE_SIZE])

	if err != nil {
		fmt.Printf("errrrrrrrrrr table")
	}

	var customer Customer

	err = customer.UnmarshalBinary(data[TABLE_SIZE:])

	if err != nil {
		fmt.Printf("errrrrrrrrrr customer")
	}

	r.Id = int(data[0])
	r.Guests = int(data[1])
	r.Table = table
	r.Customer = customer

	return nil
}
