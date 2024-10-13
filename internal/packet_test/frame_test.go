package packettest

import (
	"testing"

	"github.com/Rhisiart/MenuBridge/internal/packet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func readFrame(framer *packet.Framer) *packet.Package {
	select {
	case frame := <-framer.NewFrame():
		return frame
	}
}

func TestDecode(t *testing.T) {
	framer := packet.NewFramer()
	ch := make(chan []byte, 11)

	go framer.Frames(ch)

	ch <- []byte{packet.VERSION, 3, 0b00001010, 0x00, 0x03, 0x01, 0x02, 0x03}
	/*require.Nil(t, readFrame(framer))
	ch <- []byte{3}
	require.Nil(t, readFrame(framer))
	ch <- []byte{0b00001010}
	require.Nil(t, readFrame(framer))
	ch <- []byte{0x00, 0x03} // length 3
	require.Nil(t, readFrame(framer))
	ch <- []byte{0x01}
	require.Nil(t, readFrame(framer))
	ch <- []byte{0x02}
	require.Nil(t, readFrame(framer))
	ch <- []byte{0x03}*/
	require.Equal(t, packet.NewPackage(3, 0b1010, []byte{0x01, 0x02, 0x03}), readFrame(framer))
}

func TestEncode(t *testing.T) {
	buf := make([]byte, 12)
	f := packet.NewPackage(byte(2), byte(1), []byte("Testing"))

	f.Encode(buf, 0, 1)

	assert.Equal(t, []byte{0x1, 0x2, 0x1, 0x0, 0x7, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67}, buf)
}
