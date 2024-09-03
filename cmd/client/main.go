package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/Rhisiart/MenuBridge/internal/database"
	"github.com/Rhisiart/MenuBridge/internal/protocol"
)

func start(id int, wait *sync.WaitGroup) {
	client, err := net.Dial("tcp", fmt.Sprintf("127.0.0.%d:8080", id))

	if err != nil {
		fmt.Printf("err")
		return
	}

	defer client.Close()

	for i := 0; i < 2; i++ {
		customer := database.NewCustomer(i+1, "Martin Garrix")
		table := database.NewTable(i+1, 4)
		reservation := database.NewReservation(i+1, customer, table, 4)
		reservationBytes := reservation.MarshalBinary()

		pkg := &protocol.Package{
			Command: 0,
			Data:    reservationBytes,
		}

		time.Sleep(10000)

		b, errMarshalBinary := pkg.MarshalBinary()

		if errMarshalBinary != nil {
			fmt.Printf("err")
			return
		}

		client.Write(b)
	}

	wait.Done()
}

func main() {
	wait := &sync.WaitGroup{}

	wait.Add(2)

	go start(1, wait)
	go start(2, wait)

	wait.Wait()
}
