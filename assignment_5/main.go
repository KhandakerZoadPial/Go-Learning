package main

import (
	"fmt"
	"net/http"
)

func main() {
	// http.HandleFunc("/user_info", loggingMiddleware(getUserInfo))
	fmt.Println("Server starting on: 8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
