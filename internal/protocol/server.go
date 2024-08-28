package protocol

import (
	"fmt"
	"net"
	"sync"
)

type Sender struct {
	Pkg  *Package
	Conn *Connection
}

type Server struct {
	Sockets []Connection
	Listen  net.Listener
	Sender  chan Sender
	mutex   sync.RWMutex
}

func NewServer(port string) (*Server, error) {
	listen, err := net.Listen("tcp", port)

	if err != nil {
		return nil, err
	}

	return &Server{
		Listen:  listen,
		Sockets: make([]Connection, 0, 10),
		Sender:  make(chan Sender, 10),
		mutex:   sync.RWMutex{},
	}, nil
}

func (s *Server) Start() error {
	id := 0

	for {
		conn, err := s.Listen.Accept()
		id++

		if err != nil {
			return fmt.Errorf("error on accepting the connection: %s", err.Error())
		}

		newConn := NewConnection(conn, id)

		fmt.Printf("Connected with id %d\n", id)

		s.mutex.Lock()
		s.Sockets = append(s.Sockets, newConn)
		s.mutex.Unlock()

		go readFromConnection(s, &newConn)
	}
}

func (s *Server) Close() {
	s.Listen.Close()
}

func readFromConnection(sv *Server, conn *Connection) {
	for {
		pkg, err := conn.Next()

		if err != nil {
			fmt.Errorf("error reading the connection/package")
			break
		}

		sv.Sender <- Sender{Pkg: pkg, Conn: conn}
	}
}
