package database

import "github.com/Rhisiart/MenuBridge/types/enum"

type Order struct {
	NumberPerson int
	Table        Table
	Status       enum.OrderStatus
}

func NewOrder(numberPerson int, table Table) *Order {
	return &Order{
		NumberPerson: numberPerson,
		Table:        table,
		Status:       enum.Preparing,
	}
}

/*func (o *Order) MarshalBinary() []byte {
	q := o.Table.MarshalBinary()

	bytes := make([]byte, 0, 1+len(q))

	bytes = append(bytes, byte(o.NumberPerson))
}*/
