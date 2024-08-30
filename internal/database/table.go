package database

type Table struct {
	Id    int
	Seats int
}

func NewTable(id int, seats int) *Table {
	return &Table{
		Id:    id,
		Seats: seats,
	}
}
