package main 

import (
	"fmt"
	"net/http"
)

// Runs on hitting /health URL
func healthHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Quick Bite app is running")
}

func main(){
	http.HandleFunc("/health", healthHandler)		// Register routes before calling

	fmt.Println("Server is running on Port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting Server: ", err)
	}
}