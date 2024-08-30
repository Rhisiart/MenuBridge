package database

import "time"

type Reservation struct {
	Id       int
	Customer *Customer
	Table    *Table
	Guests   int
	DateTime string
	CreateAt string
}

func NewReservation(
	id int,
	customer *Customer,
	table *Table,
	guests int,
	dateTime string) *Reservation {
	return &Reservation{
		Id:       id,
		Customer: customer,
		Table:    table,
		Guests:   guests,
		DateTime: dateTime,
		CreateAt: time.Now().String(),
	}
}

func (r *Reservation) UnmarshalBinary(data []byte) {

}
