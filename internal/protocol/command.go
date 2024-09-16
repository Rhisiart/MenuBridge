package protocol

import (
	"github.com/Rhisiart/MenuBridge/internal/database"
)

const (
	RESERVATION = iota
	PLACE
)

var commandMap = map[string]byte{
	"reservation": RESERVATION,
	"place":       PLACE,
}

var commandMapLookup = map[byte]string{
	RESERVATION: "reservation",
	PLACE:       "place",
}

type Command struct {
	extensions map[string]byte
	size       byte
}

func CreateReservation(data []byte) database.Reservation {
	var reservation database.Reservation

	reservation.UnmarshalBinary(data)

	return reservation
}

func CreateOrder(data []byte) database.Order {
	var order database.Order

	order.UnmarshalBinary(data)

	return order
}
