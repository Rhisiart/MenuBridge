package packet

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"log/slog"

	"github.com/Rhisiart/MenuBridge/internal/entities"
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
		order := &entities.Order{}
		order.Unmarshal(p.Data)

		menus, err := service.CategoryService.FindByOrderId(ctx, order.Id)

		if err != nil {
			slog.Error(
				"Unable to getting the categories and menus",
				"Error",
				err.Error())

			return nil, false, err
		}

		data, err := json.Marshal(menus)

		return data, false, err
	case FLOOR:
		floors, err := service.FloorService.FindAll(ctx)

		if err != nil {
			slog.Error(
				"Unable to get the floor and tables",
				"Error",
				err.Error())

			return nil, false, err
		}

		data, err := json.Marshal(floors)

		return data, false, err
	case ORDER:
		slog.Warn("Received the command order...")
		order := &entities.Order{}
		orders, err := db.ReadAll(ctx, order)

		if err != nil {
			slog.Error(
				"Unable to get the orders",
				"Error",
				err.Error())

			return nil, false, err
		}

		data, err := json.Marshal(orders)

		return data, false, err
	case PLACE:
		order := &entities.Order{}

		err := json.Unmarshal(p.Data, order)

		if err != nil {
			slog.Error(
				"Unable to unmarshal the order",
				"Command",
				"Place",
				"Data",
				p.Data)

			return nil, false, err
		}

		err = order.Transaction(ctx, db)

		if err != nil {
			slog.Error("Unable make a transation to order table", "error", err.Error())

			return nil, false, nil
		}

		data, errMarshal := json.Marshal(order)

		return data, true, errMarshal
	case COMPLETE:

		return nil, false, nil
	default:
		return nil, false, nil
	}
}

func (p *Package) Types() byte {
	return p.cmd
}
