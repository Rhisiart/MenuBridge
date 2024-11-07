package database

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
