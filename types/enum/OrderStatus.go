package enum

type OrderStatus int

const (
	Preparing OrderStatus = iota
	ReadyForPickup
	Delivered
)

func (s OrderStatus) String() string {
	return [...]string{"Preparing", "Ready For Pickup", "Delivered"}[s-1]
}

func (s OrderStatus) GetIndex() int {
	return int(s)
}
