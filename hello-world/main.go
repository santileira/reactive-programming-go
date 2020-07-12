package main

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
)

func main() {
	// By default, an Observable is lazy in the sense that it
	// emits items only once a subscription is made.
	observable := rxgo.Just("Hello, World!")()

	<-observable.ForEach(
		func(v interface{}) {
			fmt.Printf("received: %v\n", v)
		},
		func(err error) {
			fmt.Printf("error: %e\n", err)
		},
		func() {
			fmt.Println("observable is closed")
		},
	)
	//ch := observable.Observe()
	//
	//item := <-ch
	//if item.Error() {
	//	fmt.Println(item.E)
	//	return
	//}
	//
	//fmt.Print(item.V)
}
