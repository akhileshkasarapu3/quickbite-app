package service

import "strings"

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

// getOrderData returns temporary in-memory order data.
var orders = []Order{
		{
			ID:              "ord_1001",
			CustomerName:    "Akhilesh",
			RestaurantID:    "res_101",
			Items:           []string{"Chicken Biryani", "Double Ka Meetha"},
			DeliveryAddress: "1417 Sage Way, Aubrey, TX",
			Status:          "placed",
		},
		{
			ID:              "ord_1002",
			CustomerName:    "Rahul",
			RestaurantID:    "res_102",
			Items:           []string{"Farmhouse Pizza"},
			DeliveryAddress: "Dallas, TX",
			Status:          "preparing",
		},
	}


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

	order := Order{
		ID:              "ord_1001",
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
	
	for _, order := range orders {
		if order.ID == id {
			return order, true
		}
	}

	return Order{}, false
}

func GetOrder() []Order {
	return orders
}