package router

import (
	"net/http"
	"github.com/akhileshkasarapu3/quickbite/internal/handler"
)

func registerHealthRoutes(mux *http.ServeMux){
	mux.HandleFunc("/api/v1/health", handler.HealthHandler)
}