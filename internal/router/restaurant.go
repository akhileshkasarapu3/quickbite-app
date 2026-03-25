package router

import (
	"net/http"
	"github.com/akhileshkasarapu3/quickbite/internal/handler"
)

func registerRestaurantRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/restaurants", handler.GetRestaurants)
	mux.HandleFunc("/api/v1/restaurant", handler.GetRestaurantByID)
	mux.HandleFunc("/api/v1/restaurants/open", handler.GetOpenRestaurants)
	mux.HandleFunc("/api/v1/restaurants/sort", handler.GetSortedRestaurants)
	mux.HandleFunc("/api/v1/restaurants/cuisine", handler.GetRestaurantsByCuisine)
}