package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world! I am learning Go Web Development.")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Server starting on: 8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
