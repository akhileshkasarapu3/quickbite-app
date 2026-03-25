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

func GetRestaurantsByCuisine(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	cuisineType := r.URL.Query().Get("type")
	if cuisineType == "" {
		response.WriteError(w, http.StatusBadRequest, "cuisine type is required")
		return 
	}

	filteredRestaurants := service.GetRestaurantsByCuisine(cuisineType)
	response.WriteSuccess(w, http.StatusOK, filteredRestaurants)
}

// GetOpenRestaurants handles GET /api/v1/restaurants/open.
func GetOpenRestaurants(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	openRestaurants := service.GetOpenRestaurants()
	response.WriteSuccess(w, http.StatusOK, openRestaurants)
}

func GetSortedRestaurants(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	sortBy := r.URL.Query().Get("by")
	if sortBy == "" {
		response.WriteError(w, http.StatusBadRequest, "sort field is required")
		return 
	}

	order := r.URL.Query().Get("order")
	if order == "" {
		response.WriteError(w, http.StatusBadRequest, "sort field is required")
		return
	}

	sortedRestaurants, validationError := service.GetSortedRestaurants(sortBy, order)
	if validationError != "" {
		response.WriteError(w, http.StatusBadRequest, validationError)
		return
	}

	response.WriteSuccess(w, http.StatusOK, sortedRestaurants)
}

// GET /api/v1/restaurants/search?type=Indian&open=true&sort=rating&order=desc
// SearchRestaurants handles combined filtering + sorting.
func SearchRestaurants(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	cuisineType := r.URL.Query().Get("type")
	openOnly := r.URL.Query().Get("open")
	sortBy := r.URL.Query().Get("sort")
	order := r.URL.Query().Get("order")

	restaurants, validationError := service.SearchRestaurants(cuisineType, openOnly, sortBy, order)
	if validationError != "" {
		response.WriteError(w, http.StatusBadRequest, validationError)
		return 
	}

	response.WriteSuccess(w, http.StatusOK, restaurants)
}