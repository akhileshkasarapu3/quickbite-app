package handler

import (
	"encoding/json"
	"net/http"

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
	var req service.CreateOrderRequest

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

// GetOrderByID handles GET /api/v1/order?id=ord_1001.
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

	// Ask service layer for the order.
	order, found := service.GetOrderByID(orderID)
	if !found {
		response.WriteError(w, http.StatusNotFound, "order not found")
		return
	}

	// Return successful response.
	response.WriteSuccess(w, http.StatusOK, order)
}

