package relay

import (
	"net/http"
	"runtime"
	"sync"

	"github.com/gorilla/websocket"
)

type Relay struct {
	port      uint16
	uuid      string
	msgs      chan []byte
	conns     chan *Connection
	mutex     sync.RWMutex
	listeners map[int]*Connection
	id        int
	send      int
}

func NewRelay(port uint16, uuid string) *Relay {
	return &Relay{
		port:      port,
		uuid:      uuid,
		msgs:      make(chan []byte, 10),
		conns:     make(chan *Connection, 10),
		mutex:     sync.RWMutex{},
		listeners: make(map[int]*Connection),
		id:        0,
		send:      runtime.NumCPU(),
	}
}

func (relay *Relay) Start() {
}

func (relay *Relay) Messages() chan []byte {
	return relay.msgs
}

func (relay *Relay) NewConnections() chan *Connection {
	return relay.conns
}

func (relay *Relay) relayRange(listeners []*Connection, data []byte, wait *sync.WaitGroup) {
	for _, conn := range listeners {
		conn.msg <- data

		wait.Done()
	}
}

func (relay *Relay) relay(data []byte) {

}

func (relay *Relay) remove(id int) {
	relay.mutex.Lock()
	delete(relay.listeners, id)
	relay.mutex.Unlock()
}

func (relay *Relay) add(id int, ws *websocket.Conn) {

}

func (relay *Relay) render(w http.ResponseWriter, r *http.Request) {
}
