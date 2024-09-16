package protocol

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"syscall"
)

type Socket struct {
	Pkg  *Package
	Conn *Connection
}

type Server struct {
	sockets []Connection
	listen  net.Listener
	Socket  chan Socket
	mutex   sync.RWMutex
}

func NewServer(port string) (*Server, error) {
	listen, err := net.Listen("tcp", port)

	if err != nil {
		return nil, err
	}

	return &Server{
		listen:  listen,
		sockets: make([]Connection, 0, 10),
		Socket:  make(chan Socket, 10),
		mutex:   sync.RWMutex{},
	}, nil
}

func (s *Server) Start() error {
	id := 0

	for {
		conn, err := s.listen.Accept()
		id++

		if err != nil {
			return fmt.Errorf("error on accepting the connection: %s", err.Error())
		}

		newConn := NewConnection(conn, id)

		fmt.Printf("Connected with id %d\n", id)

		s.mutex.Lock()
		s.sockets = append(s.sockets, newConn)
		s.mutex.Unlock()

		go readFromConnection(s, &newConn)
	}
}

func (s *Server) Send(pkg *Package) {
	s.mutex.RLock()
	removals := make([]int, 0)

	for i, conn := range s.sockets {
		err := conn.Writer.Write(pkg)

		if err != nil {
			if errors.Is(err, syscall.EPIPE) {
				fmt.Printf("connection closed by client %d\n", i)
			} else {
				fmt.Printf("removing due to error: %d, %s\n", i, err)
			}

			removals = append(removals, i)
		}
	}

	s.mutex.RUnlock()

	if len(removals) > 0 {
		s.mutex.Lock()

		for i := len(removals) - 1; i >= 0; i-- {
			idx := removals[i]
			s.sockets = append(s.sockets[:idx], s.sockets[idx+1:]...)
		}

		s.mutex.Unlock()
	}
}

func (s *Server) Close() {
	s.listen.Close()
}

func readFromConnection(sv *Server, conn *Connection) {
	defer conn.Conn.Close()

	for {
		pkg, err := conn.Next()

		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Printf("socket received EOF on connection %d\n", conn.Id)
			} else {
				fmt.Printf(
					"received error while reading from socket, on connection %d, error %s\n",
					conn.Id,
					err,
				)
			}

			break
		}

		sv.Socket <- Socket{Pkg: pkg, Conn: conn}
	}
}
