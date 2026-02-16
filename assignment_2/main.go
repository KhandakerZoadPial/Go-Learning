package main

import (
	"fmt"
	"time"
	"unicode/utf8"
)

// interfaces
type LogProvider interface {
	Log(message string) error
	Name() string
}

// structs
type ConsoleWriter struct{}
type FileWriter struct{}
type DatabaseWriter struct{}
type LogError struct {
	ProviderName string
	Timestamp    string
}

type LogStatus struct {
	ProviderName string
	Err          error
}

// errors
func (le LogError) Error() string {
	return fmt.Sprintf("Provider [%s] failed at: [%s]", le.ProviderName, le.Timestamp)
}

// interface implementation
func (cw ConsoleWriter) Name() string {
	return "Console Writer"
}
func (fw FileWriter) Name() string {
	return "File Writer"
}
func (dw DatabaseWriter) Name() string {
	return "Database Writer"
}
func (cw ConsoleWriter) Log(message string) error {
	fmt.Println("A sample log sent to terminal")
	return nil
}
func (fw FileWriter) Log(message string) error {
	time.Sleep(1 * time.Second)
	fmt.Printf("Log Message [%s] was written to file successfully\n", message)
	return nil
}
func (dw DatabaseWriter) Log(message string) error {
	fmt.Println(len(message))
	if utf8.RuneCountInString(message) > 10 {
		return LogError{ProviderName: dw.Name(), Timestamp: time.Now().String()}
	} else {
		time.Sleep(5 * time.Second)
		fmt.Printf("Log Message [%s] was written to database successfully\n", message)
		return nil
	}
}

func shipLog(p LogProvider, msg string, ch chan<- LogStatus) {
	var output = p.Log(msg)

	ch <- LogStatus{ProviderName: p.Name(), Err: output}
}

func main() {
	var providers = []LogProvider{ConsoleWriter{}, FileWriter{}, DatabaseWriter{}}
	var channel = make(chan LogStatus, len(providers))

	go shipLog(providers[0], "User Login", channel)
	go shipLog(providers[1], "User Login", channel)
	go shipLog(providers[2], "User Purchase Item 99", channel)

	for i := 0; i < len(providers); i++ {
		select {
		case output := <-channel:
			if output.Err != nil {
				fmt.Println(output.Err)
				continue
			}
			fmt.Printf("Completed Log for: [%s]\n", output.ProviderName)
		case <-time.After(2 * time.Second):
			fmt.Printf("System: ProviderTimed Out\n")
		}
	}

}
