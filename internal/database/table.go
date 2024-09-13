package database

import "github.com/Rhisiart/MenuBridge/types/enum"

type Table struct {
	Id    int
	Seats int
}

func NewTable(id int, seats int) Table {
	return Table{
		Id:    id,
		Seats: seats,
	}
}

func (t *Table) MarshalBinary() []byte {
	bytes := make([]byte, 0, 2)

	bytes = append(bytes, byte(t.Id))
	bytes = append(bytes, byte(t.Seats))

	return bytes
}

func (t *Table) UnmarshalBinary(data []byte) error {
	t.Id = int(data[0])
	t.Seats = int(data[1])

	return nil
}

func (t *Table) GetType() int {
	return enum.Table.GetIndex()
}
