package service

import (
	"sort"
	"strings"

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

// GetOpenRestaurants returns only restaurants that are currently open.
//
// This is business logic because it applies a rule top of stored restaurant data.
func GetOpenRestaurants() []model.Restaurant {
	restaurants := repository.GetRestaurants()

	var openRestaurants []model.Restaurant

	for _, restaurant := range restaurants {
		if restaurant.IsOpen == true {
			openRestaurants = append(openRestaurants, restaurant)
		}
	}
	return openRestaurants 
}


func GetRestaurantsByCuisine(cuisine string) []model.Restaurant {
	restaurants := repository.GetRestaurants()

	normalizedCuisine := strings.TrimSpace(strings.ToLower(cuisine))
	var filteredRestaurants []model.Restaurant 

	for _, restaurant := range restaurants {
		if strings.ToLower(restaurant.Cuisine) == normalizedCuisine {
			filteredRestaurants = append(filteredRestaurants, restaurant)
		} 
	}

	return filteredRestaurants
}

func GetSortedRestaurants(sortBy string, order string) ([]model.Restaurant, string){
	restaurants := repository.GetRestaurants()

	normalizedSortBy := strings.TrimSpace(strings.ToLower(sortBy))
	normalizedOrder  := strings.TrimSpace(strings.ToLower(order))

	if normalizedOrder == "" {
		normalizedOrder = "asc"
	}

	if normalizedOrder != "asc" && normalizedOrder != "desc" {
		return nil, "invalid sort order, use asc or desc"
	}

	switch normalizedSortBy {
	case "rating":
		sort.Slice(restaurants, func(i, j int)	bool {
			if normalizedOrder == "asc" {
				return restaurants[i].Rating < restaurants[j].Rating
			}
			return restaurants[i].Rating > restaurants[j].Rating
		})

	case "eta":
		sort.Slice(restaurants, func(i, j int) bool {
			if normalizedOrder == "asc" {
				return restaurants[i].ETAInMin < restaurants[j].ETAInMin
			}
			return restaurants[i].ETAInMin > restaurants[j].ETAInMin
		})

	default:
		return nil, "invalid sort field, use rating or eta"
	}

	return restaurants, ""
}