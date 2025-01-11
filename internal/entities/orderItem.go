package entities

type OrderItem struct {
	Id       int     `json:"id,omitempty"`
	Quantity int     `json:"quantity,omitempty"`
	Price    float64 `json:"price,omitempty"`
	OrderId  int     `json:"orderId,omitempty"`
	MenuId   int     `json:"menuId,omitempty"`
}
