package database

type Table struct {
	id       int `json:"id"`
	number   int `json:"number"`
	capacity int `json:"capacity"`
}

func NewTable(id int, capacity int) Table {
	return Table{
		id:       id,
		capacity: capacity,
	}
}
