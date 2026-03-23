package response

import (
	"net/http"
	"encoding/json"
)

func WriteJson(w http.ResponseWriter, statusCode int, data any){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}