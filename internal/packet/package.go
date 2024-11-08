package packet

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"log/slog"

	"github.com/Rhisiart/MenuBridge/internal/database"
	types "github.com/Rhisiart/MenuBridge/types/interface"
)

const (
	VERSION     = 1
	HEADER_SIZE = 5
)

const (
	RESERVATION = iota
	Menu
	Floor
	PLACE
	Order
	Pay
	Payment
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

func (p *Package) Execute(
	db *database.Database,
	ctx context.Context) ([]byte, bool, error) {
	switch p.Types() {
	case Menu:
		return nil, false, nil
	case Floor:
		var f database.Floor
		var floors []types.Table

		err := db.ReadAll(ctx, f, &floors)

		if err != nil {
			slog.Error(
				"Unable to getting the floor and tables",
				"Error",
				err.Error())

			return nil, false, err
		}

		data, err := json.Marshal(floors)

		return data, false, err
	default:
		return nil, false, nil
	}
}

func (p *Package) Types() byte {
	return p.cmd
}
