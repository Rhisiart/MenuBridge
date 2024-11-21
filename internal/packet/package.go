package packet

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"log/slog"

	"github.com/Rhisiart/MenuBridge/internal/database"
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
	db *database.Database,
	ctx context.Context) ([]byte, bool, error) {
	switch p.Types() {
	case MENU:
		order := &database.Order{}

		order.Unmarshal(p.Data)
		m := &database.Category{
			OrderId: order.Id,
		}

		menus, err := db.ReadAll(ctx, m)

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
		var f database.Floor

		floors, err := db.ReadAll(ctx, f)

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
		order := &database.Order{}

		order.Unmarshal(p.Data)
		err := db.Read(ctx, order)

		if err != nil {
			slog.Error(
				"Unable to get the order",
				"Error",
				err.Error())

			return nil, false, err
		}

		data, err := json.Marshal(order)

		slog.Warn("Sending data...", "data", data)

		return data, false, err
	case PLACE:
		order := &database.Order{}

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

		for _, item := range order.OrderItem {
			if item.Id != 0 {
				db.Update(ctx, item)

				slog.Warn("OrderItem", "id", item.Id, "quantity", item.Quantity, "price", item.Price)
			}
		}

		//send the order updated to all the clients
		return nil, false, nil
	default:
		return nil, false, nil
	}
}

func (p *Package) Types() byte {
	return p.cmd
}
