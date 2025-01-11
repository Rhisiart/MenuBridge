package entities

type Order struct {
	Id         int         `json:"id,omitempty"`
	CustomerId int         `json:"customerId,omitempty"`
	Amount     float64     `json:"amount,omitempty"`
	Statuscode string      `json:"statuscode,omitempty"`
	CreatedOn  string      `json:"createdOn,omitempty"`
	FloorTable FloorTable  `json:"floorTable,omitempty"`
	OrderItems []OrderItem `json:"orderItems,omitempty"`
}

func (o *Order) Unmarshal(data []byte) {
	o.Id = int(data[0])
}
