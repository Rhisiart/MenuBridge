package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/Rhisiart/MenuBridge/internal/protocol"
)

func start(id int, wait *sync.WaitGroup) {
	client, err := net.Dial("tcp", fmt.Sprintf("127.0.0.%d:8080", id))

	if err != nil {
		fmt.Printf("err")
		return
	}

	defer client.Close()

	for i := 0; i < 1000; i++ {
		pkg := &protocol.Package{
			Command: 'a',
			Data:    []byte(fmt.Sprintf("%d:%d", i, id)),
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
