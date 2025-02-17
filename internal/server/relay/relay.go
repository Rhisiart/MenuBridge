package relay

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime"
	"sync"
	"sync/atomic"

	"github.com/Rhisiart/MenuBridge/internal/server/packet"
	"github.com/gorilla/websocket"
)

type Sender struct {
	ConnId int32
	Pkg    *packet.Package
}

type Relay struct {
	port      uint16
	uuid      string
	pkgs      chan Sender
	conns     chan *Connection
	listeners map[int32]*Connection
	mutex     sync.RWMutex
	id        int32
	cores     int
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewRelay(port uint16, uuid string) *Relay {
	return &Relay{
		port:      port,
		uuid:      uuid,
		pkgs:      make(chan Sender, 10),
		conns:     make(chan *Connection, 10),
		listeners: make(map[int32]*Connection),
		mutex:     sync.RWMutex{},
		id:        0,
		cores:     runtime.NumCPU(),
	}
}

func (r *Relay) Start() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		r.render(w, req)
	})

	addr := fmt.Sprintf("0.0.0.0:%d", r.port)

	err := http.ListenAndServe(addr, nil)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
	}
}

func (r *Relay) Packages() chan Sender {
	return r.pkgs
}

func (r *Relay) NewConnections() chan *Connection {
	return r.conns
}

func (r *Relay) transport(connId int32, pkg *packet.Package) {
	select {
	case r.pkgs <- Sender{ConnId: connId, Pkg: pkg}:
	default:
	}
}

func (r *Relay) broadcastBatch(listeners []*Connection, data []byte, wait *sync.WaitGroup) {
	for _, conn := range listeners {
		conn.message(data)
	}

	wait.Done()
}

func (r *Relay) Broadcast(data []byte) {
	r.mutex.RLock()

	wait := sync.WaitGroup{}
	batchsize := len(r.listeners) / (r.cores + 1)
	batch := make([]*Connection, 0, batchsize)

	for _, listerner := range r.listeners {
		if len(batch) == batchsize {
			wait.Add(1)

			go r.broadcastBatch(batch, data, &wait)

			batch = make([]*Connection, 0, batchsize)
		}

		batch = append(batch, listerner)
	}

	if len(batch) > 0 {
		wait.Add(1)

		go r.broadcastBatch(batch, data, &wait)
	}

	wait.Wait()
	r.mutex.RUnlock()
}

func (r *Relay) Send(listenersId int32, data []byte) {
	if conn, ok := r.listeners[listenersId]; ok {
		conn.message(data)
	}
}

func (r *Relay) remove(id int32) {
	slog.Warn("Lost the connection with ", "id", id)

	r.mutex.Lock()
	delete(r.listeners, id)
	r.mutex.Unlock()
}

func (r *Relay) add(id int32, ws *websocket.Conn) {
	conn := NewConnection(id, ws, r)

	r.mutex.Lock()
	r.listeners[id] = conn
	r.mutex.Unlock()

	select {
	case r.conns <- conn:
	default:
	}

	go conn.read()
	go conn.write()
	go conn.frameReader.Frames()
	go conn.readFrames()
}

func (r *Relay) render(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)

	if err != nil {
		slog.Error("Error when render a connection", "method", "render", "message", err.Error())
		return
	}

	id := atomic.AddInt32(&r.id, 1)

	slog.Warn("New connection", "id", id)
	r.add(id, conn)
}
