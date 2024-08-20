package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	orderChan := make(chan Order, 3)

	var owg sync.WaitGroup
	go func() {
		defer owg.Done()
		owg.Add(1)
		prepareOrder(orderChan, &owg)
	}()

	// Create and send orders
	for i := 0; i < 3; i++ { // represents 3 waiters
		wg.Add(1)
		order := Order{
			Details: "rice",
			Table:   i,
		}
		go takeOrder(orderChan, order, &wg)

	}

	// Wait for all takeOrder goroutines to finish
	wg.Wait()
	close(orderChan)
	owg.Wait()

}

type Order struct {
	Details string
	Table   int
}

func takeOrder(ch chan<- Order, tableFood Order, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Taken order and informed")
	ch <- tableFood
}

func prepareOrder(ch <-chan Order, wg *sync.WaitGroup) {
	for value := range ch {
		fmt.Println("Food is prepared", value)
	}
}
