package service

import (
	"strings"
	"fmt"
	"sync"
)

// CreateOrderRequest represents the incoming order creation payload.
type CreateOrderRequest struct {
	CustomerName    string   `json:"customer_name"`
	RestaurantID    string   `json:"restaurant_id"`
	Items           []string `json:"items"`
	DeliveryAddress string   `json:"delivery_address"`
}

// Order represents the created order returned by the API.
type Order struct {
	ID              string   `json:"id"`
	CustomerName    string   `json:"customer_name"`
	RestaurantID    string   `json:"restaurant_id"`
	Items           []string `json:"items"`
	DeliveryAddress string   `json:"delivery_address"`
	Status          string   `json:"status"`
}

var (
	orders = []Order{}
	nextOrderNumber = 1001
	orderStorageLock sync.Mutex
)

// CreateOrder validates input and builds a new order.
func CreateOrder(req CreateOrderRequest) (Order, string) {
	customerName := strings.TrimSpace(req.CustomerName)
	restaurantID := strings.TrimSpace(req.RestaurantID)
	deliveryAddress := strings.TrimSpace(req.DeliveryAddress)

	if customerName == "" {
		return Order{}, "customer name is required"
	}

	if restaurantID == "" {
		return Order{}, "restaurant id is required"
	}

	if len(req.Items) == 0 {
		return Order{}, "at least one item is required"
	}

	if deliveryAddress == "" {
		return Order{}, "delivery address is required"
	}

	// Lock before touching shared writable state.
	orderStorageLock.Lock()
	defer orderStorageLock.Unlock()

	orderID := fmt.Sprintf("ord_%d", nextOrderNumber)
	nextOrderNumber++

	order := Order{
		ID:              orderID,
		CustomerName:    customerName,
		RestaurantID:    restaurantID,
		Items:           req.Items,
		DeliveryAddress: deliveryAddress,
		Status:          "placed",
	}

	orders = append(orders, order)

	return order, ""
}

// GetOrderByID searches for an order by ID.
func GetOrderByID(id string) (Order, bool) {
	// Lock while reading shared state too.
	orderStorageLock.Lock()
	defer orderStorageLock.Unlock()

	for _, order := range orders {
		if order.ID == id {
			return order, true
		}
	}

	return Order{}, false
}

func GetOrders() []Order {
	orderStorageLock.Lock()
	defer orderStorageLock.Unlock()

	// Return a copy, not the original slice.
	result := make([]Order, len(orders))
	copy(result, orders)

	return result
}