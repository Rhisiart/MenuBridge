package entities

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
