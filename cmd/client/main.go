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

	customer := database.NewCustomer(1, "Martin Garrix")
	table := database.NewTable(1, 4)
	reservation := database.NewReservation(1, customer, table, 4)
	reservationBytes := reservation.MarshalBinary()

	pkg := &protocol.Package{
		Command: 0,
		Data:    reservationBytes,
	}

	b, errMarshalBinary := pkg.MarshalBinary()

	if errMarshalBinary != nil {
		fmt.Printf("err")
		return
	}

	client.Write(b)

	time.Sleep(100000)
	wait.Done()
}

func main() {
	wait := &sync.WaitGroup{}

	wait.Add(1)

	go start(1, wait)
	//go start(2, wait)

	wait.Wait()
}
