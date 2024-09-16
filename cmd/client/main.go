package main

import (
	"fmt"
	"net"
	"sync"

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

	customer := database.NewCustomer(id, "Martin Garrix")
	table := database.NewTable(id, 4)
	reservation := database.NewReservation(id, customer, table, 4)
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
	wait.Done()
}

func main() {
	wait := sync.WaitGroup{}
	wait.Add(200)

	for i := 1; i < 201; i++ {
		go start(i, &wait)
	}

	wait.Wait()
	//start(1)
}
