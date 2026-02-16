package main

import "fmt"

type Notifier interface {
	Send(message string, address string) error
}

type Email struct{}
type SMS struct{}
type SmsError struct {
	Phone  string
	Reason string
}

func (email Email) Send(message string, address string) error {
	fmt.Printf("Email Sent: '[%s]' at [%s]\n", message, address)
	return nil
}

func (sms SMS) Send(message string, address string) error {
	if address == "" {
		return SmsError{Phone: "N/A", Reason: "You did not pass the address"}
	} else {
		fmt.Printf("SMS Sent: '[%s]' at [%s]\n", message, address)
	}
	return nil
}

func (e SmsError) Error() string {
	return fmt.Sprintf("SMS failed for [%s]: [%s]", e.Phone, e.Reason)
}

func NotifyUser(n Notifier, msg string, addr string) {
	n.Send(msg, addr)
}

func main() {
	// var e = Email{}
	var s = SMS{}

	// var notificationGateways = []Notifier{Email{}, SMS{}}

	// for _, item := range notificationGateways {
	// 	item.Send("Hello world", "fdskfjldsk")
	// }

	// NotifyUser(e, "Hello World", "test@gmail.com")
	// NotifyUser(s, "Hello World", "")
	err := s.Send("Hello World", "")

	if err != nil {
		fmt.Println(err)
	}

}
