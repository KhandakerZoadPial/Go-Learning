package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

type PriceFetcher interface {
	GetPrice(productName string) (float64, error)
	Name() string
}

type Amazon struct{}
type eBay struct{}
type Walmart struct{}

type PriceError struct {
	VendorName string
	Reason     string
}

type PriceResult struct {
	VendorName string
	Price      float64
	Err        error
}

func (pe PriceError) Error() string {
	return fmt.Sprintf("Vendor [%s] failed: %s", pe.VendorName, pe.Reason)
}

func (a Amazon) GetPrice(productName string) (float64, error) {
	time.Sleep(time.Duration(rand.IntN(5)+1) * time.Second)
	return 100.00, nil
}

func (e eBay) GetPrice(productName string) (float64, error) {
	time.Sleep(time.Duration(rand.IntN(5)+1) * time.Second)
	return 0, PriceError{VendorName: "eBay", Reason: "Product Out of Stock"}
}

func (w Walmart) GetPrice(productName string) (float64, error) {
	time.Sleep(time.Duration(rand.IntN(5)+1) * time.Second)
	return 98.5, nil
}

func (a Amazon) Name() string  { return "Amazon" }
func (e eBay) Name() string    { return "eBay" }
func (w Walmart) Name() string { return "Walmart" }

func fetchPrice(f PriceFetcher, productName string, ch chan<- PriceResult) {
	price, err := f.GetPrice(productName)

	ch <- PriceResult{
		VendorName: f.Name(),
		Price:      price,
		Err:        err,
	}
}

func main() {
	vendors := []PriceFetcher{Amazon{}, eBay{}, Walmart{}}
	channel := make(chan PriceResult, len(vendors))

	for _, item := range vendors {
		go fetchPrice(item, "Macbook Pro", channel)
	}

	for i := 0; i < len(vendors); i++ {
		select {
		case response := <-channel:
			if response.Err != nil {
				fmt.Println(response.Err)
				continue
			}
			fmt.Printf("Fetched %s: $%.2f\n", response.VendorName, response.Price)

		case <-time.After(3 * time.Second):
			fmt.Println("System: A vendor timed out.")

		}

	}
}
