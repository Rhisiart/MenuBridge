package main

import (
	"log/slog"

	"github.com/Rhisiart/MenuBridge/internal/packet"
	"github.com/Rhisiart/MenuBridge/internal/relay"
)

func main() {
	server := relay.NewRelay(8080, "123")
	framer := packet.NewFramer()

	go server.Start()
	go framer.Frames(server.Messages())

	for {
		select {
		case msg := <-server.Messages():
			slog.Warn("Message is ", "message", msg)
		case conn := <-server.NewConnections():
			slog.Warn("Established a new connection with id ", "id", conn.Id)
		}
	}
}
