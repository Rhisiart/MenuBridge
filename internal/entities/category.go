package entities

type Category struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Menus   []*Menu `json:"menus"`
	OrderId int     `json:"orderId,omitempty"`
}
