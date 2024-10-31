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

func HandleEvent(pkg *Package) ([]byte, bool, error) {
	switch pkg.Types() {
	case Menu:
		return nil, false, nil
	case Floor:
		floor := database.NewFloor(1, "Ground Floor")
		floor2 := database.NewFloor(2, "1st Floor")
		floor3 := database.NewFloor(3, "2nd Floor")
		floor4 := database.NewFloor(4, "3nd Floor")
		floor5 := database.NewFloor(5, "Rooftop")

		desk := database.NewTable(1, 1, 4)
		deskTwo := database.NewTable(2, 1, 5)
		deskThree := database.NewTable(3, 1, 6)

		floor.AddTables([]database.Table{desk, deskTwo})
		floor2.AddTables([]database.Table{desk, deskTwo, deskThree})
		floor3.AddTables([]database.Table{desk, deskTwo})
		floor4.AddTables([]database.Table{desk, deskTwo})
		floor5.AddTables([]database.Table{desk, deskTwo})

		data, err := json.Marshal([]database.Floor{floor, floor2, floor3, floor4, floor5})

		return data, false, err
	default:
		return nil, false, nil
	}
}
