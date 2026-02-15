package main


import (
	"fmt"
	"sync"
	"time"
)

type Email struct{
	Address string
	Body string
}

func (address *Email) sendEmail(wg *sync.WaitGroup){
	defer wg.Done()

	time.Sleep(3*time.Second)
	fmt.Printf("Email sent to [%s]\n", address.Address)
}


func main(){
	var wg sync.WaitGroup

	var emails = []Email{}
	emails = append(emails, Email{Address: "user1@gmail.com", Body: "This is the body"})
	emails = append(emails, Email{Address: "user2@gmail.com", Body: "This is the body"})
	emails = append(emails, Email{Address: "user3@gmail.com", Body: "This is the body"})

	fmt.Printf("Starting Batch\n")

	for _, email:= range emails{
		wg.Add(1)
		go email.sendEmail(&wg)
	}

	wg.Wait()
	fmt.Printf("Finished sending all emails.\n")

}