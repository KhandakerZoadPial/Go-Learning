package main

import "fmt"

func main() {
	var status = make(map[string]string)

	status["Auth-Service"] = "Active"
	status["Payment-Service"] = "Active"

	fmt.Println("Status before:", status)
	status["Payment-Service"] = "Down"
	fmt.Println("Status After:", status)

	// delete(status, "Auth-Service")

	authStatus, ok := status["Auth-Service"]

	if !ok {
		fmt.Println("No service found called Auth-Service")
	} else {
		fmt.Println("Auth Status is:", authStatus)
	}

	// fmt.Println("Status After Deleting Auth:", status)
}
