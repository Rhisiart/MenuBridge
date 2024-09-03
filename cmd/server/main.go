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

	fmt.Printf("-----------------------------------------------------\n")
	fmt.Printf("Reservation id: %d\n", r.Id)
	fmt.Printf("Number of guets: %d\n", r.Guests)
	fmt.Printf("Table Id: %d\n", r.Table.Id)
	fmt.Printf("Table Seats occupied: %d\n", r.Table.Seats)
	fmt.Printf("Customer Id: %d\n", r.Customer.Id)
	fmt.Printf("Customer Name: %s\n", r.Customer.Name)
}

func testingServer() {
	sv, err := protocol.NewServer(":8080")

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	defer sv.Close()
	go sv.Start()
	go sv.Hub()

	for {

	}

	/*for {
		pkg := <-sv.Socket

		fmt.Printf("-----------------------------------------------------\n")
		fmt.Printf("connection id: %d\n", pkg.Conn.Id)
		fmt.Printf("package command: %b\n", pkg.Pkg.Command)
		fmt.Printf("package data: %s\n", pkg.Pkg.Data)
	}*/
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
