package packet

import (
	"encoding/binary"
)

const (
	VERSION     = 1
	HEADER_SIZE = 5
)

type Package struct {
	cmd  byte
	Seq  byte
	Data []byte
}

func NewPackage(cmd byte, seq byte, data []byte) *Package {
	return &Package{
		cmd:  cmd,
		Seq:  seq,
		Data: data,
	}
}

func encodeHeader(data []byte, idx int, t byte, seq byte) {
	data[idx] = VERSION
	data[idx+1] = t
	data[idx+2] = seq
}

func (p *Package) Encode(idx int, seq byte) []byte {
	data := make([]byte, HEADER_SIZE+len(p.Data))

	encodeHeader(data, idx, p.Types(), seq)
	binary.BigEndian.PutUint16(data[3+idx:], uint16(len(p.Data)))
	copy(data[HEADER_SIZE+idx:], p.Data)

	return data
}

func (p *Package) Types() byte {
	return p.cmd
}
