package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// structs
type UserData struct {
	Username string `json:"username"`
	Status   string `json:"status"`
	ID       int    `json:"id"`
}

// middlewares
func loggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("LOG: %s %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next(w, r)
	}
}

// handlers
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world! I am learning Go Web Development.")
}

func getUserInfo(w http.ResponseWriter, r *http.Request) {
	var profile = UserData{
		Username: "khandakerzoadpial",
		Status:   "active",
		ID:       2,
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(profile)
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/user_info", loggingMiddleware(getUserInfo))
	fmt.Println("Server starting on: 8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
