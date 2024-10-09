package packet

import (
	"log/slog"
)

const (
	MAX_SIZE         = 1024
	MAX_PACKAGE_SIZE = 1000000
)

type Framer struct {
	previous []byte
	scratch  []byte
	frames   chan *Frame
}

func NewFramer() *Framer {
	return &Framer{
		frames: make(chan *Frame),
	}
}

func (f *Framer) packages() chan *Frame {
	return f.frames
}

func (f *Framer) Frames(frames chan []byte) {
	for {
		select {
		case frame := <-frames:
			slog.Warn("Received a frame", "message", frame)
		}
	}
}
