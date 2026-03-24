package router

import (
	"net/http"

	"github.com/akhileshkasarapu3/quickbite/internal/handler"
)

func registerOrderRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/v1/orders", func(w http.ResponseWriter, r *http.Request) {
		// Same path, different methods.
		if r.Method == http.MethodPost {
			handler.CreateOrder(w, r)
			return
		}

		if r.Method == http.MethodGet {
			handler.GetOrders(w, r)
			return
		}

		handler.GetOrders(w, r)
	})

	mux.HandleFunc("/api/v1/order", handler.GetOrderByID)
}