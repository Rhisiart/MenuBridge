package database

import (
	"sync"

	"github.com/Rhisiart/MenuBridge/types"
	"github.com/Rhisiart/MenuBridge/types/enum"
)

type Cache struct {
	mutex        sync.Mutex
	reservations []Reservation
	orders       []Order
}

func NewCache() *Cache {
	return &Cache{
		reservations: make([]Reservation, 10),
		orders:       make([]Order, 10),
	}
}

func (c *Cache) AddItem(item types.Tables) {
	c.mutex.Lock()

	switch item.GetType() {
	case enum.Reservation:
		c.reservations = append(c.reservations, item)
	}

	c.mutex.Unlock()
}
