package protocol

import (
	"encoding/binary"
	"fmt"
	"io"
)

const SCRATCH_SIZE = 1024
const MAX_PACKAGE_SIZE = 100000

type FrameReader struct {
	Reader   io.Reader
	Previous []byte
	Scratch  []byte
}

func NewFrameReader(reader io.Reader) *FrameReader {
	return &FrameReader{
		Reader:   reader,
		Previous: []byte{},
		Scratch:  make([]byte, SCRATCH_SIZE),
	}
}

func (fr *FrameReader) dataLength(frame []byte) int {
	if len(frame) < HEADER_SIZE {
		return -1
	}

	return HEADER_SIZE + int(binary.BigEndian.Uint16(frame[2:4]))
}

func (fr *FrameReader) canParse(frame []byte, dataLength int) bool {
	// Validate if the frame isn't empty, so just has the header info
	if len(frame) < HEADER_SIZE {
		return false
	}

	return len(frame) >= dataLength
}

func (fr *FrameReader) Read() ([]byte, error) {
	for {
		dataLength := fr.dataLength(fr.Previous)

		if dataLength > MAX_PACKAGE_SIZE {
			return nil, fmt.Errorf("the data length exceeded, the length define for the data %d is above %d", dataLength, MAX_PACKAGE_SIZE)
		}

		if fr.canParse(fr.Previous, dataLength) {
			out := fr.Previous[:dataLength]
			execeededDataLen := len(fr.Previous) - dataLength

			new := make([]byte, execeededDataLen)
			copy(new, fr.Previous[dataLength:])

			fr.Previous = new

			return out, nil
		}

		n, err := fr.Reader.Read(fr.Scratch)

		if err != nil {
			return nil, fmt.Errorf("error on reading the scratch")
		}

		fr.Previous = append(fr.Previous, fr.Scratch[:n]...)
	}
}
