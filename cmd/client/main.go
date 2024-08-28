package main

import (
	"fmt"
	"net"

	"github.com/Rhisiart/MenuBridge/internal/protocol"
)

func main() {
	client, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		fmt.Printf("err")
		return
	}

	defer client.Close()

	for i := 0; i < 5; i++ {
		pkg := &protocol.Package{
			Command: 'a',
			Data:    []byte("69:4201"),
		}

		b, errMarshalBinary := pkg.MarshalBinary()

		if errMarshalBinary != nil {
			fmt.Printf("err")
			return
		}

		client.Write(b)
	}
}
