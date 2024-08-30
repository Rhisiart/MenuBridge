package database

type Customer struct {
	Id   int16
	Name string
}

func NewCustomer(id int16, name string) *Customer {
	return &Customer{
		Id:   id,
		Name: name,
	}
}

func (c *Customer) MarshalBinary() ([]byte, error) {
	return nil, nil
}
