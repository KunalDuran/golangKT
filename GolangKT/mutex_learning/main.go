package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()

			// fmt.Println("incrementing by 1")
			counter++

			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(counter)
}
