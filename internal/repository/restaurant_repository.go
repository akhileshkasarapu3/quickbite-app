package repository

import "github.com/akhileshkasarapu3/quickbite/internal/model"


var restaurants = []model.Restaurant{
	{
		ID:       "res_101",
		Name:     "Hyderabad Biryani House",
		Cuisine:  "Indian",
		Rating:   4.7,
		IsOpen:   true,
		ETAInMin: 25,
	},
	{
		ID:       "res_102",
		Name:     "Pizza Point",
		Cuisine:  "Italian",
		Rating:   4.3,
		IsOpen:   true,
		ETAInMin: 18,
	},
	{
		ID:       "res_103",
		Name:     "Sushi Express",
		Cuisine:  "Japanese",
		Rating:   4.5,
		IsOpen:   false,
		ETAInMin: 0,
	},
}

func GetRestaurants() []model.Restaurant {
	result := make([]model.Restaurant, len(restaurants))
	copy(result, restaurants)

	return result
}

func GetRestaurantByID(id string) (model.Restaurant, bool) {
	for _, restaurant := range restaurants {
		if restaurant.ID == id {
			return restaurant, true
		}
	}

	return model.Restaurant{}, false
} 