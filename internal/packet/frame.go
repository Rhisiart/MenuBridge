package packet

import (
	"fmt"

	"github.com/Rhisiart/MenuBridge/internal/utils"
)

const (
	VERSION     = 1
	HEADER_SIZE = 5
)

type Frame struct {
	cmd  byte
	seq  byte
	data []byte
}

func NewFrame(cmd byte, seq byte, data []byte) *Frame {
	return &Frame{
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

func (f *Frame) Encode(data []byte, idx int, seq byte) (int, error) {
	if len(data) > HEADER_SIZE+len(f.data) {
		return -1, fmt.Errorf("the buffer doent have enough sace for the frame header and data")
	}

	EncodeHeader(data, idx, f.Types(), seq)
	utils.Write16(data, 3+idx, len(f.data))
	copy(data[HEADER_SIZE+idx:], f.data)

	return HEADER_SIZE + len(f.data), nil
}

func (f *Frame) Types() byte {
	return f.cmd
}
