package packettest

import (
	"testing"

	"github.com/Rhisiart/MenuBridge/internal/packet"
	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	buf := make([]byte, 12)
	f := packet.NewFrame(byte(2), byte(1), []byte("Testing"))

	f.Encode(buf, 0, 1)

	assert.Equal(t, []byte{0x1, 0x2, 0x1, 0x0, 0x7, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67}, buf)
}
