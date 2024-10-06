package relay

import (
	"log/slog"

	"github.com/gorilla/websocket"
)

type Connection struct {
	Id    int32
	conn  *websocket.Conn
	relay *Relay
	msg   chan []byte
}

func NewConnection(id int32, conn *websocket.Conn, relay *Relay) *Connection {
	return &Connection{
		Id:    id,
		conn:  conn,
		relay: relay,
		msg:   make(chan []byte, 10),
	}
}

func (c *Connection) read() {
	for {
		msgType, data, err := c.conn.ReadMessage()

		slog.Warn("reading from the message on connection ", "id", c.Id, "message", string(data))

		if err != nil {
			slog.Warn("error", "method", "read", "error", err.Error())
			break
		}

		if msgType != websocket.BinaryMessage {
			slog.Warn("The message type instead binary")
			break
		}

		c.relay.broadcast(data)
	}

	slog.Warn("Closing the connection", "id", c.Id, "method", "read")

	c.relay.remove(c.Id)
	c.conn.Close()
}

func (c *Connection) write() {
	for {
		msg := <-c.msg

		err := c.conn.WriteMessage(websocket.BinaryMessage, msg)

		slog.Warn("writting from the connection", "id", c.Id, "message", string(msg))

		if err != nil {
			slog.Warn("error", "method", "write", "error", err.Error())
			break
		}
	}

	slog.Warn("Closing the connection", "id", c.Id, "method", "write")

	c.relay.remove(c.Id)
	c.conn.Close()
}

func (c *Connection) message(data []byte) {
	select {
	case c.msg <- data:
	default:
	}
}
