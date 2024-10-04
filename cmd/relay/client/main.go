package main

import (
	"log/slog"

	"github.com/Rhisiart/MenuBridge/internal/relay"
)

type frame struct {
	uuid string
	msg  []byte
}

func KeepClientAlive(host string, path string, uuid string, msgs chan frame) {
	client := relay.NewRelayDriver(host, path, uuid)

	err := client.Connect()

	if err != nil {
		slog.Error(err.Error())
		return
	}

	go func() {
		defer client.Close()

		for {
			_, data, _ := client.Conn.ReadMessage()

			msgs <- frame{
				uuid: uuid,
				msg:  data}
		}
	}()

	return
}

func main() {
	msgs := make(chan frame, 10)

	KeepClientAlive("localhost:8080", "ws", "1", msgs)
	//KeepClientAlive("localhost:8080", "ws", "2", msgs)
	//KeepClientAlive("localhost:8080", "ws", "3", msgs)

	writter := relay.NewRelayDriver("localhost:8080", "ws", "4")

	defer writter.Close()

	messages := []string{
		"This is very difficul",
		"Im trying my best",
		"Ill understand",
	}

	for _, message := range messages {
		writter.Relay([]byte(message))

		for range 1 {
			select {
			case msg := <-msgs:
				slog.Warn("Received message ", "Message", msg.msg, "id", msg.uuid)
			}
		}
	}
}
