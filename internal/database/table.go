package database

type Table struct {
	Id       int `json:"id,omitempty"`
	Number   int `json:"number,omitempty"`
	Capacity int `json:"capacity,omitempty"`
}

func NewTable(id int, capacity int) Table {
	return Table{
		Id:       id,
		Capacity: capacity,
	}
}

func (t *Table) Unmarshal(data []byte) {

}
