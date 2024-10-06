package main

import (
	"log/slog"
	"time"

	"github.com/Rhisiart/MenuBridge/internal/relay"
)

type frame struct {
	uuid string
	msg  []byte
}

func KeepClientAlive(host string, path string, uuid string, msgs chan frame) {
	client := relay.NewRelayDriver(host, path, uuid)

	err := client.Connect()
	defer client.Close()

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
	msgs := make(chan frame, 250)

	KeepClientAlive("localhost:8080", "ws", "1", msgs)
	//KeepClientAlive("localhost:8080", "ws", "2", msgs)
	//KeepClientAlive("localhost:8080", "ws", "3", msgs)

	writter := relay.NewRelayDriver("localhost:8080", "ws", "4")
	writter.Connect()

	defer writter.Close()
	<-time.NewTimer(time.Millisecond * 500).C

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
				slog.Warn("Received message ", "Message", string(msg.msg), "id", msg.uuid)
			case <-time.NewTimer(time.Second).C:
				slog.Warn("waiting for message", "line", message)
			}
		}
	}
}
