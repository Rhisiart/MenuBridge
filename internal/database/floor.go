package database

type Floor struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Tables []Table `json:"tables"`
}

func NewFloor(id int, name string) Floor {
	return Floor{
		Id:   id,
		Name: name,
	}
}

func (f *Floor) AddTables(tables []Table) {
	f.Tables = tables
}

func (f *Floor) MarshalBinary() []byte {
	b := make([]byte, 0, 1+len(f.Name))

	b = append(b, byte(f.Id))
	b = append(b, f.Name...)

	return b
}

func (f *Floor) UnmarshalBinary(data []byte) {
	f.Id = int(data[0])
	f.Name = string(data[1:])
}
