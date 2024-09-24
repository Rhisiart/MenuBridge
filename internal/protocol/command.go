package protocol

import (
	"github.com/Rhisiart/MenuBridge/internal/database"
)

const (
	RESERVATION = iota
	PLACE
	Order
	Pay
	Payment
)

var commandMap = map[string]byte{
	"reservation": RESERVATION,
	"place":       PLACE,
	"order":       Order,
	"Pay":         Pay,
	"Payment":     Payment,
}

var commandMapLookup = map[byte]string{
	RESERVATION: "reservation",
	PLACE:       "place",
	Order:       "order",
	Pay:         "Pay",
	Payment:     "Payment",
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

func GetOrder(data []byte) database.Order {
	var order database.Order

	order.UnmarshalBinary(data)

	return order
}

func MakeOrderItem(data []byte) database.OrderItem {
	var orderItem database.OrderItem

	orderItem.UnmarshalBinary(data)

	return orderItem
}

func SendPayment(data []byte) {

}
