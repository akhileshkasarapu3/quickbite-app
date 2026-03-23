package main 

import (
	"fmt"
	"net/http"
	"github.com/akhileshkasarapu3/quickbite/internal/router"
)



func main(){
	// Register all application routes.
	appRouter := router.RegisterRoutes()

	fmt.Println("Server is running on Port 8080")

	// Start the HTTP server using our router.
	err := http.ListenAndServe(":8080", appRouter)
	if err != nil {
		fmt.Println("Error starting Server: ", err)
	}
}