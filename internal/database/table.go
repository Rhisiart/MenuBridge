package database

type Table struct {
	id       int
	capacity int
}

func NewTable(id int, capacity int) Table {
	return Table{
		id:       id,
		capacity: capacity,
	}
}

func (t *Table) MarshalBinary() []byte {
	bytes := make([]byte, 0, 2)

	bytes = append(bytes, byte(t.id))
	bytes = append(bytes, byte(t.capacity))

	return bytes
}

func (t *Table) UnmarshalBinary(data []byte) error {
	t.id = int(data[0])
	t.capacity = int(data[1])

	return nil
}
