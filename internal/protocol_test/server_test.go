package protocol_test

import (
	"net"
	"testing"

	"github.com/Rhisiart/MenuBridge/internal/protocol"
)

func TestServer(t *testing.T) {
	client, err := net.Dial("tcp", "127.0.0.1:8080")

	if err != nil {
		t.FailNow()
	}

	defer client.Close()

	pkg := &protocol.Package{
		Command: 'a',
		Data:    []byte("69:4201"),
	}

	b, errMarshalBinary := pkg.MarshalBinary()

	if errMarshalBinary != nil {
		t.FailNow()
	}

	client.Write(b)
}
