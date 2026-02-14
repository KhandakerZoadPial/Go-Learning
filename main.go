package main

import (
	"errors"
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func getUser(id int) (string, error) {
	if id == 42 {
		return "admin", nil
	} else {
		return "", errors.New("user not found")
	}
}

func main() {

	a := 10
	b := 5

	additionResult := add(a, b)
	fmt.Println("Addition result is:", additionResult)

	name, error := getUser(100)

	if error != nil {
		fmt.Println("Error:", error)
	} else {
		fmt.Println("The users name is:", name)
	}
}
