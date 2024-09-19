package main

import (
	"fmt"
	"net"

	"github.com/Rhisiart/MenuBridge/internal/database"
	"github.com/Rhisiart/MenuBridge/internal/protocol"
)

func start(id int) {
	client, err := net.Dial("tcp", fmt.Sprintf("127.0.0.%d:8080", id))

	if err != nil {
		fmt.Printf("err")
		return
	}

	defer client.Close()

	//menu := database.NewMenu(id, "Meat", "Meat", 4)
	customer := database.NewCustomer(id, "Martin Garrix")
	table := database.NewTable(id, 4)
	order := database.NewOrder(id, 2, table, customer)
	//orderLine := database.NewOrderItem(id, menu, order, 1)
	reservation := database.NewReservation(id, customer, table, 4)

	SendPackage(client, &protocol.Package{
		Command: 0,
		Data:    reservation.MarshalBinary(),
	})

	SendPackage(client, &protocol.Package{
		Command: 1,
		Data:    order.MarshalBinary(),
	})
}

func SendPackage(client net.Conn, pkg *protocol.Package) {
	b, errMarshalBinary := pkg.MarshalBinary()

	if errMarshalBinary != nil {
		fmt.Printf("err")
		return
	}

	client.Write(b)
}

func main() {

	start(1)
	/*wait := sync.WaitGroup{}
	wait.Add(200)

	for i := 1; i < 201; i++ {
		go start(i, &wait)
	}

	wait.Wait()*/
	//start(1)
}
