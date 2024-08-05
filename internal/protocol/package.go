package protocol

import (
	"encoding/binary"
	"fmt"
)

const HEADER_SIZE = 4
const VERSION = 1

type Package struct {
	Command byte
	Data    []byte
}

func (p *Package) Marshal() ([]byte, error) {
	ulen := binary.BigEndian.Uint16(p.Data)

	dataLength := make([]byte, 2)
	pack := make([]byte, HEADER_SIZE+ulen)

	binary.BigEndian.PutUint16(dataLength, ulen)

	pack = append(pack, VERSION)
	pack = append(pack, p.Command)
	pack = append(pack, dataLength...)
	pack = append(pack, p.Data...)

	return pack, nil
}

func (p *Package) Unmarshal(bytes []byte) error {
	if VERSION != bytes[0] {
		return fmt.Errorf("unable to unmarshal package, because the versions doesnt match")
	}

	validLength := int(binary.BigEndian.Uint16(bytes[2:4]))
	end := HEADER_SIZE + validLength
	packageLen := len(bytes)

	if packageLen <= end {
		return fmt.Errorf("invalid length of data, the data has %d should be equal or above %d", packageLen, end)
	}

	p.Command = bytes[1]
	p.Data = bytes[HEADER_SIZE:end]

	return nil
}
