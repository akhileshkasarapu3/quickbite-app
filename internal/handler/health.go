package handler

import (
	"net/http"
	"encoding/json"
)

type HealthResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struc{
	Status string `json:"status"`
	Message string `json:"message"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		errorResponse := ErrorResponse{
			Status: "error"
			Message: "method not allowed",
		}

		err := json.NewEncoder(w).Encode(errResponse)
		if err != nil {
			http.Error(w, "Failed to Encode Response", http.StatusInternalServerError)
			return
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)		// Set http status code 200

	response := HealthResponse{
		Status : "Success",
		Message : "Quick Bite is running",
	}

	// converts the Go map into JSON and writes it into the response.
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to Encode response", http.StatusInternalServerError)
		return 
	}
}