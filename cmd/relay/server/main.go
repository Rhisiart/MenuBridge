package main

import (
	"log/slog"
	"os"
	"runtime"
	"strconv"

	"github.com/Rhisiart/MenuBridge/internal/packet"
	"github.com/Rhisiart/MenuBridge/internal/relay"
)

func main() {
	runtime.GOMAXPROCS(runtime.GOMAXPROCS(0) - 1)

	var port uint = 0

	if port == 0 {
		portStr := os.Getenv("PORT")
		portEnv, err := strconv.Atoi(portStr)
		if err == nil {
			port = uint(portEnv)
		}
	}

	uuid := os.Getenv("AUTH_ID")

	slog.Warn("port selected", "port", port)
	r := relay.NewRelay(uint16(port), uuid)

	go onMessage(r)

	r.Start()
}

/*func newConnections(relay *relay.Relay) {
	for {
		conn := <-relay.NewConnections()

	}
}*/

func onMessage(relay *relay.Relay) {
	framer := packet.NewFramer()
	go framer.Frames(relay.Messages())

	for {
		select {
		case frame := <-framer.NewFrame():
			slog.Warn("received frame", "frame", frame)
		}
	}
}
