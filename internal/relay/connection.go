package relay

import "github.com/gorilla/websocket"

type Connection struct {
	id    int
	conn  *websocket.Conn
	relay *Relay
	msg   chan []byte
}

func NewConnection(id int, conn *websocket.Conn, relay *Relay) *Connection {
	return &Connection{
		id:    id,
		conn:  conn,
		relay: relay,
		msg:   make(chan []byte, 10),
	}
}
