package main

import (
	"context"
	"log/slog"
	"runtime"

	"github.com/Rhisiart/MenuBridge/internal/config"
	"github.com/Rhisiart/MenuBridge/internal/database"
	"github.com/Rhisiart/MenuBridge/internal/packet"
	"github.com/Rhisiart/MenuBridge/internal/relay"
)

func main() {
	runtime.GOMAXPROCS(runtime.GOMAXPROCS(0) - 1)

	ctx := context.Background()
	config, err := config.NewConfiguration()
	uuid := "1"

	if err != nil {
		slog.Error("Unable to get the environment variables", "Error", err.Error())
		return
	}

	db := database.NewDatabase(config.DatabaseUrl)
	err = db.Connect()
	defer db.Close()

	if err != nil {
		slog.Error("Unable to connect to the database", "Error", err.Error())
		return
	}

	slog.Warn("port selected", "port", config.Port)
	r := relay.NewRelay(uint16(config.Port), uuid)

	go onMessage(r, db, ctx)
	go newConnections(r)

	r.Start()
}

func newConnections(relay *relay.Relay) {
	for {
		conn := <-relay.NewConnections()

		slog.Warn("New connection", "Id", conn.Id)
	}
}

func onMessage(relay *relay.Relay, db *database.Database, ctx context.Context) {
	for {
		frame := <-relay.Packages()

		slog.Warn("received frame", "Connection", frame.ConnId, "frame", frame.Pkg.Data)

		data, broadcast, err := packet.HandleEvent(db, ctx, frame.Pkg)

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
