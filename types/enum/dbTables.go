package enum

type DbTables int

const (
	Table DbTables = iota
	Reservation
	Order
	Customer
)

func (dbt DbTables) String() string {
	return [...]string{"Table", "Reservation", "Order", "Customer"}[dbt-1]
}

func (dbt DbTables) GetIndex() int {
	return int(dbt)
}
