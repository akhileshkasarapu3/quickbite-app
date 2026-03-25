package service

import (
	"github.com/akhileshkasarapu3/quickbite/internal/model"
	"github.com/akhileshkasarapu3/quickbite/internal/repository"
)

// GetRestaurants asks repository for all restaurants.
func GetRestaurants() []model.Restaurant {
	return repository.GetRestaurants()
}

// GetRestaurantByID asks repository for one restaurant by ID.
func GetRestaurantByID(id string) (model.Restaurant, bool) {
	return repository.GetRestaurantByID(id)
}