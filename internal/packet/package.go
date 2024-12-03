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
		f := &database.Floor{}

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
		slog.Warn("Received the command order...")
		order := &database.Order{}
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
		order := &database.Order{}
		broadcast := true

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

		err = db.Transaction(ctx, order)

		if err != nil {
			slog.Error("Unable make a transation to order table", "error", err.Error())

			return nil, false, nil
		}

		/*for _, item := range order.OrderItem {
			var err error

			if item.Id != 0 {
				err = db.Update(ctx, item)
			} else {
				item.OrderId = order.Id
				err = db.Create(ctx, item)
			}

			if err != nil {
				slog.Error(
					"Unable to update the order item",
					"id",
					item.Id,
					"quantity",
					item.Quantity,
					"price",
					item.Price)
			} else {
				slog.Warn("OrderItem", "id", item.Id, "quantity", item.Quantity, "price", item.Price)
			}
		}*/

		data, errMarshal := json.Marshal(order)

		//send the order updated to all the clients
		return data, broadcast, errMarshal
	default:
		return nil, false, nil
	}
}

func (p *Package) Types() byte {
	return p.cmd
}
