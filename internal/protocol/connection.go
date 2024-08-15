package protocol

import "net"

type Connection struct {
	FrameReader *FrameReader
	FrameWriter *FrameWriter
	Id          int
	Conn        net.Conn
}

func NewConnection(conn net.Conn, id int) *Connection {
	return &Connection{
		FrameReader: NewFrameReader(conn),
		FrameWriter: NewFrameWriter(conn),
		Id:          id,
		Conn:        conn,
	}
}

func (conn *Connection) Next() (*Package, error) {
	frame, err := conn.FrameReader.Read()

	if err != nil {
		return nil, err
	}

	var p Package
	unmarshalErr := p.Unmarshal(frame)

	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return &p, nil
}
