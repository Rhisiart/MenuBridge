package relay

import (
	"github.com/gorilla/websocket"
)

type Connection struct {
	id    int32
	conn  *websocket.Conn
	relay *Relay
	msg   chan []byte
}

func NewConnection(id int32, conn *websocket.Conn, relay *Relay) *Connection {
	return &Connection{
		id:    id,
		conn:  conn,
		relay: relay,
		msg:   make(chan []byte, 10),
	}
}

func (c *Connection) Read() {
	for {
		msgType, data, err := c.conn.ReadMessage()

		if err != nil || msgType != websocket.CloseMessage {
			break
		}

		c.relay.Broadcast(data)
	}

	c.relay.remove(c.id)
	c.conn.Close()
}

func (c *Connection) Write() {
	for {
		msg := <-c.msg

		err := c.conn.WriteMessage(websocket.BinaryMessage, msg)

		if err != nil {
			break
		}
	}

	c.relay.remove(c.id)
	c.conn.Close()
}

func (c *Connection) Message(data []byte) {
	select {
	case c.msg <- data:
	default:
	}
}
