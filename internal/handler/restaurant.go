package handler

import (
	"net/http"
	"github.com/akhileshkasarapu3/quickbite/internal/response"
)


// One restaurant
type Restaurant struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Cuisine string `json:"cuisine"`
	Rating float64 `json:"rating"`
	IsOpen bool	`json:"is_open"`
	ETAInMin int `json:"eta_in_min"`
}

func GetRestaurants(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return 
	}

	restaurants := []Restaurant {
		{
			ID:   "res_001",
			Name:	"Spice Rack",
			Cuisine: "Indian",
			Rating: 4.0,
			IsOpen: true,
			ETAInMin: 25,
		},
		{
			ID:   "res_002",
			Name:	"Pizza Point",
			Cuisine: "Indian",
			Rating: 4.2,
			IsOpen: false,
			ETAInMin: 15,
		},
	}

	response.WriteSuccess(w, http.StatusOK, restaurants)
}