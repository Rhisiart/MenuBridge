package database

type FloorTable struct {
	Id      int `json:"id,omitempty"`
	Number  int `json:"number,omitempty"`
	TableId int `json:"tableId,omitempty"`
	FloorId int `json:"floorId,omitempty"`
}
