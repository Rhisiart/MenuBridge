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
	buf    []byte
	Data   chan []byte
	frames chan *Package
}

func NewFramer() *Framer {
	return &Framer{
		buf:    make([]byte, 0),
		Data:   make(chan []byte, 10),
		frames: make(chan *Package),
	}
}

func (f *Framer) decode() error {
	if f.buf[0] != VERSION {
		return fmt.Errorf("the version received %d dont match with %d version", f.buf[0], VERSION)
	}

	if len(f.buf) < HEADER_SIZE {
		return nil
	}

	dataLen := int(binary.BigEndian.Uint16(f.buf[3:5]))
	totalLength := (HEADER_SIZE + dataLen)
	exceededBytes := len(f.buf) - totalLength

	if len(f.buf) < totalLength {
		return nil
	}

	f.frames <- NewPackage(f.buf[1], f.buf[2], f.buf[HEADER_SIZE:totalLength])
	copy(f.buf, f.buf[totalLength:])
	f.buf = f.buf[:exceededBytes]

	return nil
}

func (f *Framer) NewFrame() chan *Package {
	return f.frames
}

func (f *Framer) Frames() {
	for {
		data := <-f.Data
		slog.Warn("Receiving a frame", "data", data, "Buffer", len(f.buf))

		f.buf = append(f.buf, data...)

		if len(f.buf) > HEADER_SIZE {
			err := f.decode()

			if err != nil {
				slog.Error("fail to decode the package", "message", err.Error())
			}
		}
	}
}
