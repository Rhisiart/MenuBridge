package main

import (
	"fmt"
	"time"
)

var ch chan int

func read() {
	for {
		select {
		case msg := <-ch:
			fmt.Println("Received from ch1:", msg)
		}
	}
}

func main() {
	ch := make(chan int)

	//go read()

	go func() {
		time.Sleep(time.Second * 1)
		ch <- 1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ch <- 2
	}()

	for {
		fmt.Println(<-ch)
	}
}
