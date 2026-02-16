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

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	var newUser UserData
	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Created User: %s\n", newUser.Username)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func main() {
	http.HandleFunc("/add_user", loggingMiddleware(createUser))

	// server stuff
	fmt.Println("Server starting on: 8080...")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server: ", err)
	}
}
