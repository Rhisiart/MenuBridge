package main

import (
	"log/slog"
	"time"

	"github.com/Rhisiart/MenuBridge/internal/server/packet"
	"github.com/Rhisiart/MenuBridge/internal/server/relay"
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

func runRelayClients() {
	msgs := make(chan frame, 250)

	KeepClientAlive("localhost:8080", "ws", "1", msgs)
	KeepClientAlive("localhost:8080", "ws", "2", msgs)
	KeepClientAlive("localhost:8080", "ws", "3", msgs)

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

		for range 3 {
			select {
			case msg := <-msgs:
				slog.Warn("Received message ", "Message", string(msg.msg), "id", msg.uuid)
			case <-time.NewTimer(time.Second).C:
				slog.Warn("waiting for message", "line", message)
			}
		}
	}
}

func main() {
	writter := relay.NewRelayDriver("localhost:8080", "ws", "4")
	writter.Connect()

	defer writter.Close()

	for i := 0; i < 5; i++ {

		pkg := packet.NewPackage(byte(2), byte(i+1), []byte{byte(i), 0x01, 0x02})
		data := pkg.Encode(0, byte(i+1))

		writter.Relay(data)

		time.Sleep(10 * time.Second)
	}
}
