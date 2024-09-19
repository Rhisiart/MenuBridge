package protocol

import (
	"github.com/Rhisiart/MenuBridge/internal/database"
)

const (
	RESERVATION = iota
	PLACE
	Order
)

var commandMap = map[string]byte{
	"reservation": RESERVATION,
	"place":       PLACE,
	"order":       Order,
}

var commandMapLookup = map[byte]string{
	RESERVATION: "reservation",
	PLACE:       "place",
	Order:       "order",
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

func MakeOrderItem(data []byte) database.OrderItem {
	var orderItem database.OrderItem

	orderItem.UnmarshalBinary(data)

	return orderItem
}
