package service

import (
	"strings"

	"github.com/akhileshkasarapu3/quickbite/internal/model"
	"github.com/akhileshkasarapu3/quickbite/internal/repository"
)

func CreateOrder(req model.CreateOrderRequest) (model.Order, string) {
	// Transform data
	customerName := strings.TrimSpace(req.CustomerName)
	restaurantID := strings.TrimSpace(req.RestaurantID)
	deliveryAddress := strings.TrimSpace(req.DeliveryAddress)

	// validate the data
	if customerName == "" {
		return model.Order{}, "customer name is required"
	}
	if restaurantID == "" {
		return model.Order{}, "restaurant id is required"
	}
	if len(req.Items) == 0 {
		return model.Order{}, "at least one item is required"
	}
	if deliveryAddress == "" {
		return model.Order{}, "delivery address is required"
	}

	// build object
	order := model.Order{
		CustomerName:    customerName,
		RestaurantID:    restaurantID,
		Items:           req.Items,
		DeliveryAddress: deliveryAddress,
		Status:          "placed",
	}

	// Ask repository to store it and generate ID.
	savedOrder := repository.SaveOrder(order)

	return savedOrder, ""
}

func GetOrderByID(id string) (model.Order, bool) {
	// Use it from repo layer.
	return repository.GetOrderByID(id)
}

func GetOrders() []model.Order {
	return repository.GetOrders()
}