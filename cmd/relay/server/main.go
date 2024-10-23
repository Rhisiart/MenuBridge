package main

import (
	"encoding/json"
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
	go newConnections(r)

	r.Start()
}

func newConnections(relay *relay.Relay) {
	for {
		conn := <-relay.NewConnections()

		slog.Warn("New connection", "Id", conn.Id)
	}
}

func onMessage(relay *relay.Relay) {
	framer := packet.NewFramer()
	go framer.Frames(relay.Messages())

	for {
		frame := <-framer.NewFrame()
		slog.Warn("received frame", "frame", frame)

		switch frame.Types() {
		case 2:
			slog.Warn("Sending the menus", "Command", 2)

			jsonData, err := json.Marshal([]string{"bitoque", "bacalhau com natas"})

			if err != nil {
				slog.Error("Couldnt marshal the menus", "error", err.Error())
				return
			}

			data := make([]byte, 37)
			pkg := packet.NewPackage(byte(2), byte(1), jsonData)
			_, errEncode := pkg.Encode(data, 0, byte(1))

			if errEncode != nil {
				slog.Error("Couldnt encode the package", "error", errEncode.Error())
				return
			}

			relay.Send(1, data)
		}
	}
}
