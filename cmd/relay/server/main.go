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
	for {
		frame := <-relay.Packages()

		slog.Warn("received frame", "Connection", frame.ConnId, "frame", frame.Pkg.Data)

		data, broadcast, err := packet.HandleEvent(frame.Pkg)

		if err != nil {
			slog.Error(
				"Could not handle the event",
				"Command",
				frame.Pkg.Types(),
				"err",
				err.Error())
		}

		pkg := packet.NewPackage(frame.Pkg.Types(), byte(1), data)
		pkgEncoded := pkg.Encode(0, byte(1))

		if broadcast {
			relay.Broadcast(pkgEncoded)
		} else {
			relay.Send(frame.ConnId, pkgEncoded)
		}

	}
}
