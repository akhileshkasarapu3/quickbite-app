package router

import (
	"net/http"

	"github.com/akhileshkasarapu3/quickbite/internal/handler"
)

func registerOrderRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/orders", handler.CreateOrder)
}