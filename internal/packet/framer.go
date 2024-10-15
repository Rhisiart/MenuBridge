package packet

import (
	"encoding/binary"
	"fmt"
	"log/slog"
)

const (
	MAX_SIZE         = 1024
	MAX_PACKAGE_SIZE = 1000000
)

type Framer struct {
	data   []byte
	frames chan *Package
}

func NewFramer() *Framer {
	return &Framer{
		data:   make([]byte, 0),
		frames: make(chan *Package),
	}
}

func (f *Framer) decode() error {
	if f.data[0] != VERSION {
		return fmt.Errorf("the version received %d dont match with %d version", f.data[0], VERSION)
	}

	if len(f.data) < HEADER_SIZE {
		return nil
	}

	dataLen := int(binary.BigEndian.Uint16(f.data[3:5]))
	totalLength := (HEADER_SIZE + dataLen)
	exceededBytes := len(f.data) - totalLength

	if len(f.data) < totalLength {
		return nil
	}

	f.frames <- NewPackage(f.data[1], f.data[2], f.data[HEADER_SIZE:totalLength])
	copy(f.data, f.data[totalLength:])
	f.data = f.data[:exceededBytes]

	return nil
}

func (f *Framer) NewFrame() chan *Package {
	return f.frames
}

func (f *Framer) Frames(data chan []byte) {
	for {
		if len(f.data) > HEADER_SIZE {
			err := f.decode()

			if err != nil {
				slog.Error("fail to decode the package", "message", err.Error())
			}
		}

		f.data = append(f.data, <-data...)
	}
}
