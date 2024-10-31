package database

type Floor struct {
	id   int
	name string
}

func NewFloor(id int, name string) Floor {
	return Floor{
		id:   id,
		name: name,
	}
}

func (f *Floor) MarshalBinary() []byte {
	b := make([]byte, 0, 1+len(f.name))

	b = append(b, byte(f.id))
	b = append(b, f.name...)

	return b
}

func (f *Floor) UnmarshalBinary(data []byte) {
	f.id = int(data[0])
	f.name = string(data[1:])
}
