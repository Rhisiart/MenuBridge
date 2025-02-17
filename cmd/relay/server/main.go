package main

import (
	"context"
	"log/slog"
	"runtime"

	"github.com/Rhisiart/MenuBridge/internal/config"
	"github.com/Rhisiart/MenuBridge/internal/server/packet"
	"github.com/Rhisiart/MenuBridge/internal/server/relay"
	"github.com/Rhisiart/MenuBridge/internal/service"
	"github.com/Rhisiart/MenuBridge/internal/storage"
	"github.com/Rhisiart/MenuBridge/internal/storage/postgres"
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

	db := postgres.NewDatabase(config.DatabaseUrl)
	err = db.Connect()
	defer db.Close()

	if err != nil {
		slog.Error("Unable to connect to the database", "Error", err.Error())
		return
	}

	repository := storage.NewRepository(db.Database)
	service := service.NewService(repository)

	slog.Warn("port selected", "port", config.Port)
	r := relay.NewRelay(uint16(config.Port), uuid)

	go onMessage(r, service, ctx)
	go newConnections(r)

	r.Start()
}

func newConnections(relay *relay.Relay) {
	for {
		conn := <-relay.NewConnections()

		slog.Warn("New connection", "Id", conn.Id)
	}
}

func onMessage(relay *relay.Relay, service *service.Service, ctx context.Context) {
	for {
		frame := <-relay.Packages()

		slog.Warn("received frame", "Connection", frame.ConnId, "frame", frame.Pkg.Data)

		data, broadcast, err := frame.Pkg.Execute(service, ctx)

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
