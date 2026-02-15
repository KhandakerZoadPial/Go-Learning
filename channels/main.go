package main

import (
	"fmt"
	"time"
)


func checkStatus(serviceName string, ch chan string){
	time.Sleep(3*time.Second)

	var message = fmt.Sprintf("Service %s is is UP", serviceName)

	ch <- message
}


func main(){
	var ch = make(chan string)

	services := []string{"Auth", "Payment", "Inventory"}

	for _,service:= range services{
		go checkStatus(service, ch)
	}

	for i:=0; i<len(services); i++{
		var result = <-ch
		fmt.Println("Result:", result)
	}
}