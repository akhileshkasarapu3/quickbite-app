package model

// CreateOrderRequest represent incoming JSON to create a new order
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
