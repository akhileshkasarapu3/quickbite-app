package handler

import (
	"net/http"

	"github.com/akhileshkasarapu3/quickbite/internal/response"
)

type HealthData struct {
	Message string `json:"message"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "method not allowed")
	}

	healthCheckData := HealthData{
		Message: "Quick Bite API is Running",
	}
	response.WriteSuccess(w, http.StatusOK, healthCheckData)
}