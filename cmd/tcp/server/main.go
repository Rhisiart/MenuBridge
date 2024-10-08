package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/Rhisiart/MenuBridge/internal/database"
	"github.com/Rhisiart/MenuBridge/internal/protocol"
)

func main() {
	//testingPackage()
	//testingPackageAndFrameReader()
	//testingFrameWriter()
	testingServer()
}

func testingMarshalAndUnMarshalOfReservation() {
	var r database.Reservation

	customer := database.NewCustomer(2, "Martin Garrix")
	table := database.NewTable(3, 5)
	reservation := database.NewReservation(1, customer, table, 4)
	reservationBytes := reservation.MarshalBinary()

	for _, b := range reservationBytes {
		fmt.Printf("%x ", b)
	}

	r.UnmarshalBinary(reservationBytes)

	/*fmt.Printf("-----------------------------------------------------\n")
	fmt.Printf("Reservation id: %d\n", r.Id)
	fmt.Printf("Number of guets: %d\n", r.Guests)
	fmt.Printf("Table Id: %d\n", r.Table.Id)
	fmt.Printf("Table Seats occupied: %d\n", r.Table.)
	fmt.Printf("Customer Id: %d\n", r.Customer.Id)
	fmt.Printf("Customer Name: %s\n", r.Customer.Name)*/
}

func testingMarshalOfMenu() {
	menu := database.NewMenu(1, "testing", "testingtesting", 20)
	m := menu.MarshalBinary()
	var mn database.Menu
	mn.UnmarshalBinary(m)
	fmt.Printf("Id = %d\n", mn.Id)
	fmt.Printf("Name = %s\n", mn.Name)
	fmt.Printf("Description = %s\n", mn.Description)
	fmt.Printf("Description len = %d\n", len(mn.Description))
	fmt.Printf("Price = %d\n", mn.Price)
}

func testingServer() {
	sv, err := protocol.NewServer(":8080")

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	cache := database.NewCache()

	defer sv.Close()
	go sv.Start()

	for {
		socket := <-sv.Socket

		fmt.Printf("---------------------------------------------------\n")
		fmt.Printf("connection id: %d\n", socket.Conn.Id)
		fmt.Printf("package command: %d\n", int(socket.Pkg.Command))

		switch socket.Pkg.Command {
		case protocol.RESERVATION:
			reservation := protocol.CreateReservation(socket.Pkg.Data)

			cache.AddItem(reservation)
			sv.Send(socket.Pkg)
		case protocol.PLACE:
			order := protocol.GetOrder(socket.Pkg.Data)

			cache.AddItem(order)
			sv.Send(socket.Pkg)
		case protocol.Order:
			orderItem := protocol.MakeOrderItem(socket.Pkg.Data)

			cache.AddItem(orderItem)
			sv.Send(socket.Pkg)
		case protocol.Pay:
			order := protocol.GetOrder(socket.Pkg.Data)
			amount := cache.CalculateOrderAmount(order.Id)
			payment := database.NewPayment(1, order, amount)

			pkg := &protocol.Package{
				Command: 4,
				Data:    payment.MarshalBinary(),
			}

			sv.Send(pkg)
		}
	}
}

func testingFrameWriter() {
	p := &protocol.Package{
		Command: 'a',
		Data:    []byte("69:4201"),
	}

	writer := bytes.NewBuffer(nil)

	fw := protocol.NewFrameWriter(writer)
	fw.Write(p)

	reader := bytes.NewReader(writer.Bytes())
	fr := protocol.NewFrameReader(reader)

	b, err := fr.Read()

	if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else {
		dataLength := int(binary.BigEndian.Uint16(b[2:4]))
		p.UnmarshalBinary(b)

		fmt.Printf("Length of data = %d \n", dataLength)
		fmt.Printf("%s \n", p.Data)
	}
}

func testingPackageAndFrameReader() {
	p := &protocol.Package{
		Command: 'a',
		Data:    []byte("69:4201"),
	}

	packageCompress, err := p.MarshalBinary()
	dLen := int(binary.BigEndian.Uint16(packageCompress[2:4]))
	fmt.Printf("package length = %d\n", dLen)

	if err != nil {
		fmt.Println("error")
		return
	}

	var packages []byte
	for i := 0; i < 100; i++ {
		packages = append(packages, packageCompress...)
	}

	reader := bytes.NewReader(packages)
	fr := protocol.NewFrameReader(reader)

	for i := 0; i < 100; i++ {
		var p protocol.Package
		fmt.Printf("Starting Reading \n")
		b, err := fr.Read()

		if err != nil {
			fmt.Printf("%s\n", err.Error())
		} else {
			dataLength := int(binary.BigEndian.Uint16(b[2:4]))
			p.UnmarshalBinary(b)

			fmt.Printf("Length of data = %d \n", dataLength)
			fmt.Printf("%s \n", p.Data)
		}

		fmt.Printf("Finished Reading the %d frame\n", i+1)
	}

}

func testingPackage() {
	p := &protocol.Package{
		Command: 'a',
		Data:    []byte("69:4201"),
	}

	length := binary.BigEndian.Uint16(p.Data)
	lengthData := make([]byte, 2)

	binary.BigEndian.PutUint16(lengthData, length)

	b := make([]byte, 0, 1+1+2+length)
	b = append(b, 1)
	b = append(b, p.Command)
	b = append(b, lengthData...)
	b = append(b, p.Data...)

	reader := bytes.NewReader(b)
	n, _ := reader.Read(b)

	fmt.Printf("%d\n", length)
	fmt.Printf("%d\n", len(b))
	fmt.Printf("%d\n", n)
	fmt.Printf("%x\n", b[:n])
	fmt.Printf("%x\n", int16(binary.BigEndian.Uint16(b[2:])))

}
