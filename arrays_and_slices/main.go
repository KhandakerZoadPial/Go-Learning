package main

import "fmt"

func main() {
	todos := []string{}

	todos = append(todos, "Learn Arrays")
	todos = append(todos, "Learn Slices")
	todos = append(todos, "Master Go")

	fmt.Println("Length of the slice is:", len(todos))

	fmt.Println("Todos before removing items:", todos)

	todos = append(todos[:1], todos[2:]...)
	fmt.Println("Todos after removing items:", todos)

}
