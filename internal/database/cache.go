package database

import (
	"sync"
)

type Cache struct {
	mutex        sync.Mutex
	reservations []Reservation
	orders       []Order
	orderItems   []OrderItem
}

func NewCache() *Cache {
	return &Cache{
		reservations: make([]Reservation, 0, 10),
		orders:       make([]Order, 0, 10),
		orderItems:   make([]OrderItem, 0, 10),
	}
}

func (c *Cache) AddItem(item interface{}) {
	c.mutex.Lock()

	switch i := item.(type) {
	case Reservation:
		c.reservations = append(c.reservations, i)
		i.Print()
	case Order:
		c.orders = append(c.orders, i)
		i.Print()
	case OrderItem:
		c.orderItems = append(c.orderItems, i)
	}

	c.mutex.Unlock()
}
