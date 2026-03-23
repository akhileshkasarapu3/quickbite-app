package handler

import (
	"net/http"

	"github.com/akhileshkasarapu3/quickbite/internal/response"
)

type HealthResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		errorResponse := ErrorResponse{
			Status: "error",
			Message: "method not allowed",
		}

		response.WriteJson(w, http.StatusMethodNotAllowed, errorResponse)
		return
	}

	successResponse := HealthResponse{
		Status : "Success",
		Message : "Quick Bite is running",
	}

	response.WriteJson(w, http.StatusOK, successResponse)
}