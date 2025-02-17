package entities

type Table struct {
	Id       int   `json:"id,omitempty"` //Id for the floor_dinnertable table
	Number   int   `json:"number,omitempty"`
	Capacity int   `json:"capacity,omitempty"`
	Order    Order `json:"order,omitempty"`
}

func NewTable(id int, capacity int) Table {
	return Table{
		Id:       id,
		Capacity: capacity,
	}
}
