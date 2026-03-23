package main 

import (
	"fmt"
	"net/http"

	"github.com/akhileshkasarapu3/quickbite/internal/handler"
)



func main(){
	http.HandleFunc("/health", handler.HealthHandler)		// Register routes before calling

	fmt.Println("Server is running on Port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting Server: ", err)
	}
}