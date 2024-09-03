package protocol

import (
	"fmt"

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

func CreateReservation(data []byte) {
	var reservation database.Reservation

	reservation.UnmarshalBinary(data)

	fmt.Printf("Reservation id: %d\n", reservation.Id)
	fmt.Printf("Table Id: %d\n", reservation.Table.Id)
	fmt.Printf("Number of guets: %d\n", reservation.Guests)
	fmt.Printf("Customer Id: %d\n", reservation.Customer.Id)
	fmt.Printf("Customer Name: %s\n", reservation.Customer.Name)
}
