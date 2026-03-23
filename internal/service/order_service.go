package service

import "strings"

type CreateOrderRequest struct {
	CustomerName    string   `json:"customer_name"`
	RestaurantID    string   `json:"restaurant_id"`
	Items           []string `json:"items"`
	DeliveryAddress string   `json:"delivery_address"`	
}

type Order struct {
	ID              string   `json:"id"`
	CustomerName    string   `json:"customer_name"`
	RestaurantID    string   `json:"restaurant_id"`
	Items           []string `json:"items"`
	DeliveryAddress string   `json:"delivery_address"`
	Status          string   `json:"status"`
}


func CreateOrder(req CreateOrderRequest) (Order, string) {
	// Trim spaces so input like "   " is treated as empty.
	customerName := strings.TrimSpace(req.CustomerName)
	restaurantID := strings.TrimSpace(req.RestaurantID)
	deliveryAddress := strings.TrimSpace(req.DeliveryAddress)

	// Validate required fields.
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

	order := Order{
		ID:              "ord_1001",
		CustomerName:    customerName,
		RestaurantID:    restaurantID,
		Items:           req.Items,
		DeliveryAddress: deliveryAddress,
		Status:          "placed",
	}

	return order, ""
}
