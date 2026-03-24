package handler

import (
	"encoding/json"
	"net/http"

	"github.com/akhileshkasarapu3/quickbite/internal/model"
	"github.com/akhileshkasarapu3/quickbite/internal/response"
	"github.com/akhileshkasarapu3/quickbite/internal/service"
)

// CreateOrder handles POST /api/v1/orders.
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Decode incoming JSON body into request struct.
	var req model.CreateOrderRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	// Ask service layer to validate and create order.
	order, validationError := service.CreateOrder(req)
	if validationError != "" {
		response.WriteError(w, http.StatusBadRequest, validationError)
		return
	}

	// Return 201 Created when a resource is created successfully.
	response.WriteSuccess(w, http.StatusCreated, order)
}

func GetOrders(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "invalid method")
		return
	}

	orders := service.GetOrders()

	response.WriteSuccess(w, http.StatusOK, orders)
}

func GetOrderByID(w http.ResponseWriter, r *http.Request) {
	// Allow only GET requests.
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// Read order ID from query string.
	orderID := r.URL.Query().Get("id")
	if orderID == "" {
		response.WriteError(w, http.StatusBadRequest, "order id is required")
		return
	}

	order, found := service.GetOrderByID(orderID)
	if !found {
		response.WriteError(w, http.StatusNotFound, "order not found")
		return
	}

	response.WriteSuccess(w, http.StatusOK, order)
}