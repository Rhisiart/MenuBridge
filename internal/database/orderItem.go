package database

type OrderItem struct {
	Id       int     `json:"id"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}
