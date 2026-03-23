package service


type Restaurant struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Cuisine  string  `json:"cuisine"`
	Rating   float64 `json:"rating"`
	IsOpen   bool    `json:"is_open"`
	ETAInMin int     `json:"eta_in_min"`
}

func GetRestaurants() []Restaurant {
	return []Restaurant{
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
}

func GetRestaurantByID(id string) (Restaurant, bool) {
	restaurants := GetRestaurants()		

	for _, restaurant := range restaurants {
		if restaurant.ID == id {
			return restaurant, true
		}
	}

	return Restaurant{}, false	// return empty interface
}