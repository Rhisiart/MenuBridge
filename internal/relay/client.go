package relay

import (
	"net/url"

	"github.com/gorilla/websocket"
)

type RelayDriver struct {
	url  url.URL
	uuid string
	conn *websocket.Conn
}

func NewRelayDriver(host string, path string, uuid string) *RelayDriver {
	u := url.URL{Scheme: "ws", Host: host, Path: path}

	return &RelayDriver{
		url:  u,
		uuid: uuid,
	}
}

func (r *RelayDriver) Connect() error {
	c, _, err := websocket.DefaultDialer.Dial(r.url.String(), nil)
	if err != nil {
		return err
	}

	r.conn = c
	return c.WriteMessage(websocket.BinaryMessage, []byte(r.uuid))
}

func (r *RelayDriver) Relay(data []byte) error {
	return r.conn.WriteMessage(websocket.BinaryMessage, data)
}

func (r *RelayDriver) Close() {
	r.conn.Close()
}
