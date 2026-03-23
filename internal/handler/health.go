package handler

import (
	"net/http"
	"encoding/json"
)

func HealthHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)		// Set http status code 200

	response := map[string]string {
		"status": "success",
		"message": "Quickbite is Running Successfully!",
	}

	// converts the Go map into JSON and writes it into the response.
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to Encode response", http.StatusInternalServerError)
	}
}