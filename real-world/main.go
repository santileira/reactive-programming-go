package main

import "github.com/reactivex/rxgo/v2"

type Customer struct {
	ID             int
	Name, LastName string
	Age            int
	TaxNumber      string
}

func main() {
	// Crete the input channel
	ch := make(chan rxgo.Item)

	// Data producer
	go producer(ch)

	// Create an Observable
	observable := rxgo.FromChannel(ch)
}
