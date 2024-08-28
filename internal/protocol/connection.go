package protocol

import "net"

type Connection struct {
	Reader *FrameReader
	Writer *FrameWriter
	Id     int
	Conn   net.Conn
}

func NewConnection(conn net.Conn, id int) Connection {
	return Connection{
		Reader: NewFrameReader(conn),
		Writer: NewFrameWriter(conn),
		Id:     id,
		Conn:   conn,
	}
}

func (conn *Connection) Next() (*Package, error) {
	frame, err := conn.Reader.Read()

	if err != nil {
		return nil, err
	}

	var p Package
	unmarshalErr := p.UnmarshalBinary(frame)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &p, nil
}
