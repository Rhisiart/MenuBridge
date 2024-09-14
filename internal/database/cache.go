package database

import (
	"sync"
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

func (c *Cache) AddItem(item interface{}) {
	c.mutex.Lock()

	switch i := item.(type) {
	case Reservation:
		c.reservations = append(c.reservations, i)
	case Order:
		c.orders = append(c.orders, i)
	}

	c.mutex.Unlock()
}
