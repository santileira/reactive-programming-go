package main

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"strconv"
)

type Customer struct {
	ID             int
	Name, LastName string
	Age            int
	TaxNumber      string
}

func main() {
	// Create the input channel
	ch := make(chan rxgo.Item)

	// Data producer
	go producer(ch)

	// Create an Observable
	observable := rxgo.FromChannel(ch)

	observable.
		Filter(
			func(item interface{}) bool {
				// Filter operation
				customer := item.(Customer)
				return customer.Age > 18
			}).
		Map(
			func(_ context.Context, item interface{}) (interface{}, error) {
				// Enrich operation
				customer := item.(Customer)
				customer.TaxNumber = getTaxNumber(customer)
				return customer, nil
			},
			// Create multiple instances of the map operator
			rxgo.WithPool(10),
			rxgo.WithBufferedChannel(1),
		)

	for customer := range observable.Observe() {
		if customer.Error() {
			continue
		}
		fmt.Println(customer)
	}
}

func getTaxNumber(customer Customer) string {
	return strconv.Itoa(customer.ID)
}

func producer(ch chan rxgo.Item) {
	i := 0
	for {

		ch <- rxgo.Item{
			V: Customer{
				ID: i,
			},
		}
		i++
	}
}
