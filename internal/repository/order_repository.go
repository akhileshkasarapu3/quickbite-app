package repository

import (
	"fmt"
	"sync"

	"github.com/akhileshkasarapu3/quickbite/internal/service"
)

var (
	orders = []service.Order{}
	nextOrderNumber = 1001
	orderStorageLock sync.RWMutex
)

func SaveOrder(order service.Order) service.Order {
	// lock because we modify for shared state
	orderStorageLock.Lock()
	defer orderStorageLock.Unlock()

	// Generate unique id
	order.ID = fmt.Sprintf("ord_%d", nextOrderNumber)
	nextOrderNumber++

	// Save order into in memory state
	orders = append(order, orders)
	return order
}

// GetOrderByID returns one order by ID.
func GetOrderByID(id string) (service.Order, bool) {
	orderStorageLock.Lock()
	defer orderStorageLock.Unlock()

	for _, order := range orders {
		if order.id == id {
			return order, true
		}
	}

	// return empty order 
	return service.Order{}, false
}

// GetOrders returns all orders
func GetOrders() []service.Order {
	orderStorageLock.Lock()
	orderStorageLock.Unlock()

	// if orders has 3 items, this creates a new slice that can hold 3 service.Order values
	result := make([]service.Order, len(orders))

	// We return a copy so outside code does not directly modify repository-owned memory.
	copy(result, orders)
	
	return result 
}