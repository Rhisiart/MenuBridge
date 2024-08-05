package main

import (
	"encoding/binary"
	"fmt"

	"github.com/Rhisiart/MenuBridge/internal/protocol"
)

func main() {
	p := &protocol.Package{
		Command: 'a',
		Data:    []byte("69:4201"),
	}

	length := uint16(len(p.Data))
	lengthData := make([]byte, 2)

	binary.BigEndian.PutUint16(lengthData, length)

	b := make([]byte, 0, 1+1+2+length)
	b = append(b, 1)
	b = append(b, p.Command)
	b = append(b, lengthData...)

	final := append(b, p.Data...)

	fmt.Printf("%d\n", length)
	fmt.Printf("%d\n", len(final))
	//fmt.Printf("%x\n", final[2:])
	fmt.Printf("%x\n", lengthData)
	fmt.Printf("%x\n", int(binary.BigEndian.Uint16(final[2:])))
}
