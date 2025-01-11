package packet

import (
	"context"
	"encoding/binary"

	"github.com/Rhisiart/MenuBridge/internal/service"
)

const (
	VERSION     = 1
	HEADER_SIZE = 5
)

const (
	OPEN = iota
	MENU
	FLOOR
	RESERVATION
	PLACE
	ORDER
	COMPLETE
	PAY
	PAYMENT
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
	service *service.Service,
	ctx context.Context) ([]byte, bool, error) {
	switch p.Types() {
	case MENU:
		data, err := service.CategoryService.FindByOrderId(ctx, p.Data)
		return data, false, err
	case FLOOR:
		data, err := service.FloorService.FindAll(ctx)
		return data, false, err
	case ORDER:
		data, err := service.OrderService.Create(ctx)
		return data, false, err
	case PLACE:
		data, err := service.OrderService.UpsertWithOrderItems(ctx, p.Data)
		return data, true, err
	case COMPLETE:

		return nil, false, nil
	default:
		return nil, false, nil
	}
}

func (p *Package) Types() byte {
	return p.cmd
}
