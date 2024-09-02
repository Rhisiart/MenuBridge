package protocol

import (
	"fmt"
	"net"
	"sync"

	"github.com/Rhisiart/MenuBridge/internal/database"
)

type Socket struct {
	Pkg  *Package
	Conn *Connection
}

type Server struct {
	Sockets []Connection
	Listen  net.Listener
	Socket  chan Socket
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
		Socket:  make(chan Socket, 10),
		mutex:   sync.RWMutex{},
	}, nil
}

func (s *Server) Start() error {
	id := 0

	go s.Hub()

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

func (s *Server) Hub() {
	for {
		select {
		case socket := <-s.Socket:
			switch socket.Pkg.Command {
			case RESERVATION:
				s.handleReservation(socket)
			}
		}
	}
}

func (s *Server) handleReservation(socket Socket) {
	var reservation database.Reservation

	reservation.UnmarshalBinary(socket.Pkg.Data)

	fmt.Printf("-----------------------------------------------------\n")
	fmt.Printf("connection id: %d\n", socket.Conn.Id)
	fmt.Printf("package command: %b\n", socket.Pkg.Command)
	fmt.Printf("Reservation id: %d\n", reservation.Id)
	fmt.Printf("Table Id: %d\n", reservation.Table.Id)
	fmt.Printf("Customer Id: %d\n", reservation.Customer.Id)
	fmt.Printf("Customer Name: %s\n", reservation.Customer.Name)
	fmt.Printf("Number of guets: %d\n", reservation.Guests)
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

		sv.Socket <- Socket{Pkg: pkg, Conn: conn}
	}
}
