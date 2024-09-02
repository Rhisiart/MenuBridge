package database

type Customer struct {
	Id   int
	Name string
}

func NewCustomer(id int, name string) Customer {
	return Customer{
		Id:   id,
		Name: name,
	}
}

// Take in might could use copy for encoding the string
// bytes[0] = byte(c.Id)
// copy(bytes[1:], c.Name)
func (c *Customer) MarshalBinary() []byte {
	bytes := make([]byte, 0, 1+len(c.Name))

	bytes = append(bytes, byte(c.Id))
	bytes = append(bytes, []byte(c.Name)...)

	return bytes
}

func (c *Customer) UnmarshalBinary(data []byte) error {
	c.Id = int(data[0])
	c.Name = string(data[1:])

	return nil
}
