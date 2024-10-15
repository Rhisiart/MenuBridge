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

		if err != nil {
			slog.Error("error", "method", "read", "error", err.Error())
			break
		}

		if msgType != websocket.BinaryMessage {
			slog.Error("The message type instead binary", "method", "read", "error", err.Error())
			break
		}

		c.relay.broadcast(data)
	}

	c.relay.remove(c.Id)
	c.conn.Close()
}

func (c *Connection) write() {
	for {
		msg := <-c.msg

		err := c.conn.WriteMessage(websocket.BinaryMessage, msg)

		if err != nil {
			slog.Error("error", "method", "write", "error", err.Error())
			break
		}
	}

	c.relay.remove(c.Id)
	c.conn.Close()
}

func (c *Connection) message(data []byte) {
	select {
	case c.msg <- data:
	default:
	}
}
