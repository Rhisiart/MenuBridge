package relay

import (
	"log/slog"

	"github.com/Rhisiart/MenuBridge/internal/server/packet"
	"github.com/gorilla/websocket"
)

type Connection struct {
	Id          int32
	conn        *websocket.Conn
	relay       *Relay
	frameReader *packet.Framer
	msg         chan []byte
}

func NewConnection(id int32, conn *websocket.Conn, relay *Relay) *Connection {
	return &Connection{
		Id:          id,
		conn:        conn,
		relay:       relay,
		frameReader: packet.NewFramer(),
		msg:         make(chan []byte, 10),
	}
}

func (c *Connection) readFrames() {
	for {
		frame := <-c.frameReader.NewFrame()

		c.relay.transport(c.Id, frame)
	}
}

func (c *Connection) read() {
	defer func() {
		c.relay.remove(c.Id)
		c.conn.Close()
	}()

	for {
		msgType, data, err := c.conn.ReadMessage()

		slog.Warn("Receiver a new message", "message", data)

		if err != nil {
			slog.Error("error", "method", "read", "error", err.Error())
			break
		}

		if msgType != websocket.BinaryMessage {
			slog.Error("The message type is not binary", "method", "read")
			break
		}

		c.frameReader.Data <- data
	}
}

func (c *Connection) write() {
	defer func() {
		c.relay.remove(c.Id)
		c.conn.Close()
	}()

	for {
		msg, ok := <-c.msg

		if !ok {
			slog.Error("Closing the channel or no more data", "method", "write", "id", c.Id)
			c.conn.WriteMessage(websocket.CloseMessage, []byte{})
			break
		}

		err := c.conn.WriteMessage(websocket.BinaryMessage, msg)

		if err != nil {
			slog.Error("Writting error", "method", "write", "error", err.Error())
			break
		}
	}
}

func (c *Connection) message(data []byte) {
	select {
	case c.msg <- data:
	default:
	}
}
