package handler

import (
	"net/http"
	"github.com/akhileshkasarapu3/quickbite/internal/response"
	"github.com/akhileshkasarapu3/quickbite/internal/service"
)


func GetRestaurants(w http.ResponseWriter, r *http.Request){

	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return 
	}

	restaurants := service.GetRestaurants() 
	response.WriteSuccess(w, http.StatusOK, restaurants)
}


func GetRestaurantByID(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return 
	}

	restaurantID := r.URL.Query().Get("id")
	if restaurantID == "" {
		response.WriteError(w, http.StatusBadRequest, "restaurant id is required")
		return
	}

	restaurant, found := service.GetRestaurantByID(restaurantID)
	if !found {
		response.WriteError(w, http.StatusNotFound, "restaurant not found")
		return
	}

	response.WriteSuccess(w, http.StatusOK, restaurant)
}