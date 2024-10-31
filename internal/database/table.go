package database

type Table struct {
	Id       int `json:"id"`
	floorId  int
	Capacity int `json:"capacity"`
}

func NewTable(id int, floorId int, capacity int) Table {
	return Table{
		Id:       id,
		floorId:  floorId,
		Capacity: capacity,
	}
}

func (t *Table) MarshalBinary() []byte {
	bytes := make([]byte, 0, 2)

	bytes = append(bytes, byte(t.Id))
	bytes = append(bytes, byte(t.Capacity))

	return bytes
}

func (t *Table) UnmarshalBinary(data []byte) error {
	t.Id = int(data[0])
	t.Capacity = int(data[1])

	return nil
}
