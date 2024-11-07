package packet

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/Rhisiart/MenuBridge/internal/database"
	types "github.com/Rhisiart/MenuBridge/types/interface"
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

func HandleEvent(
	db *database.Database,
	ctx context.Context,
	pkg *Package) ([]byte, bool, error) {
	switch pkg.Types() {
	case Menu:
		return nil, false, nil
	case Floor:
		var f database.Floor
		var floors []types.Table

		err := db.ReadAll(ctx, f, floors)

		if err != nil {
			slog.Error(
				"Unable to getting the floor and tables",
				"Error",
				err.Error())

			return nil, false, err
		}

		data, err := json.Marshal(floors)

		slog.Warn("The data is", "floors", data)

		return data, false, err
	default:
		return nil, false, nil
	}
}
