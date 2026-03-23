package router

import (
	"net/http"
)


func RegisterRoutes() http.Handler{
	// Centralize routing
		// Instead of pulling all routes in main
	
	mux := http.NewServeMux()	// 1. local switchboard
	
	// 2. Register routes 
	// version v1 api routes 
	// mux.HandleFunc("/api/v1/health", handler.HealthHandler)
	registerHealthRoutes(mux)
	registerRestaurantRoutes(mux) 

	return mux 
}