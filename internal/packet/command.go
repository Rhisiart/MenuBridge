package packet

import (
	"encoding/json"

	"github.com/Rhisiart/MenuBridge/internal/database"
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

type FloorCache struct {
	database.Floor
	Tables []database.Table
}

func HandleEvent(pkg *Package) ([]byte, bool, error) {
	switch pkg.Types() {
	case Menu:
		return nil, false, nil
	case Floor:
		f := database.NewFloor(1, "Main Floor")
		desk := database.NewTable(1, 1, 4)
		deskTwo := database.NewTable(2, 1, 4)

		floors := FloorCache{
			Floor:  f,
			Tables: []database.Table{desk, deskTwo},
		}

		data, err := json.Marshal(floors)

		return data, false, err
	default:
		return nil, false, nil
	}
}
