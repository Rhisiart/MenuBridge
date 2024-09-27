package relay

import (
	"fmt"
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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

func (relay *Relay) BroadcastBatch(listeners []*Connection, data []byte, wait *sync.WaitGroup) {
	for _, conn := range listeners {
		conn.Message(data)
	}

	wait.Done()
}

func (relay *Relay) Broadcast(data []byte) {
	select {
	case relay.msgs <- data:
	default:
	}

	relay.mutex.RLock()

	wait := sync.WaitGroup{}
	batchsize := len(relay.listeners) / (relay.send + 1)
	batch := make([]*Connection, 0, batchsize)

	for _, listerner := range relay.listeners {
		if len(batch) == batchsize {
			wait.Add(1)

			go relay.BroadcastBatch(batch, data, &wait)

			batch = make([]*Connection, 0, batchsize)
		}

		batch = append(batch, listerner)
	}

	if len(batch) > 0 {
		wait.Add(1)

		go relay.BroadcastBatch(batch, data, &wait)
	}

	wait.Wait()
	relay.mutex.RUnlock()
}

func (relay *Relay) remove(id int) {
	relay.mutex.Lock()
	delete(relay.listeners, id)
	relay.mutex.Unlock()
}

func (relay *Relay) add(id int, ws *websocket.Conn) {
	conn := NewConnection(id, ws, relay)

	relay.mutex.Lock()
	relay.listeners[id] = conn
	relay.mutex.Unlock()

	select {
	case relay.conns <- conn:
	default:
	}

	go conn.Read()
	go conn.Write()
}

func (relay *Relay) render(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("Error when render a connection")
		return
	}

	relay.mutex.Lock()
	relay.id++
	id := relay.id
	relay.mutex.Unlock()

	relay.add(id, conn)
}
