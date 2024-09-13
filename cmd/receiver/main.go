package main

import (
	"fmt"
	"net"

	"github.com/Rhisiart/MenuBridge/internal/database"
	"github.com/Rhisiart/MenuBridge/internal/protocol"
)

func main() {
	client, err := net.Dial("tcp", fmt.Sprintf("127.0.0.%d:8080", 2))

	if err != nil {
		fmt.Printf("err")
		return
	}

	for {
		response := make([]byte, 1024)
		_, err = client.Read(response)

		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		var pkg protocol.Package
		var reservation database.Reservation

		pkg.UnmarshalBinary(response)
		reservation.UnmarshalBinary(pkg.Data)

		fmt.Printf("command %b\n", pkg.Command)
		fmt.Printf("Reservation id %d\n", reservation.Id)
		fmt.Printf("Guest number: %d\n", reservation.Guests)
	}
}
