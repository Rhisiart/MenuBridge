package entities

type Floor struct {
	Id     int     `json:"id"`
	Name   string  `json:"name"`
	Tables []Table `json:"tables,omitempty"`
}
