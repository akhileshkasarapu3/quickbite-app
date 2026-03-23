package response

import (
	"net/http"
	"encoding/json"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

type APIResponse struct {
	Status string `json:"status"`
	Data 	any 	`json:"data"`
}


func WriteJson(w http.ResponseWriter, statusCode int, data any){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}
}

func WriteSuccess(w http.ResponseWriter, statusCode int, data any){
	successResponse := APIResponse{
		Status: "success",
		Data: data,
	}

	WriteJson(w, statusCode, successResponse)
}

func WriteError(w http.ResponseWriter, statusCode int, message string){
	errorResponse := ErrorResponse{
		Status: "error",
		Message: message,
	}

	WriteJson(w, statusCode, errorResponse)
}