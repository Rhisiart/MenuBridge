package packet

import (
	"encoding/binary"
	"fmt"
)

const (
	VERSION     = 1
	HEADER_SIZE = 5
)

type Package struct {
	cmd  byte
	seq  byte
	data []byte
}

func NewPackage(cmd byte, seq byte, data []byte) *Package {
	return &Package{
		cmd:  cmd,
		seq:  seq,
		data: data,
	}
}

func EncodeHeader(data []byte, idx int, t byte, seq byte) {
	data[idx] = VERSION
	data[idx+1] = t
	data[idx+2] = seq
}

func (p *Package) Encode(data []byte, idx int, seq byte) (int, error) {
	if len(data) > HEADER_SIZE+len(p.data) {
		return -1, fmt.Errorf("the buffer doent have enough sace for the frame header and data")
	}

	EncodeHeader(data, idx, p.Types(), seq)
	binary.BigEndian.PutUint16(data[3+idx:], uint16(len(p.data)))
	copy(data[HEADER_SIZE+idx:], p.data)

	return HEADER_SIZE + len(p.data), nil
}

func (p *Package) Types() byte {
	return p.cmd
}
