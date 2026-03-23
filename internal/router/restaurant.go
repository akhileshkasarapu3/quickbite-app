package router

import (
	"net/http"
	"github.com/akhileshkasarapu3/quickbite/internal/handler"
)

func registerRestaurantRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/restaurants", handler.GetRestaurants)
}