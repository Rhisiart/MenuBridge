package database

type Payment struct {
	id    int
	order Order
	//method
	amount int
}

func NewPayment(id int, order Order, amount int) Payment {
	return Payment{
		id:     id,
		order:  order,
		amount: amount,
	}
}

func (p *Payment) MarshalBinary() []byte {
	order := p.order.MarshalBinary()

	b := make([]byte, 1+len(order)+1)
	b = append(b, byte(p.id))
	b = append(b, order...)
	b = append(b, byte(p.amount))

	return b
}

func (p *Payment) UnmarshalBinary(data []byte) error {
	return nil
}
