package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Mark this goroutine as done when it completes
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main_() {
	// var wg sync.WaitGroup

	// for i := 1; i <= 3; i++ {
	// 	wg.Add(1) // Add a new worker to the WaitGroup
	// 	go worker(i, &wg)
	// }

	// // Wait for all workers to complete
	// wg.Wait()
	// fmt.Println("All workers finished.")

	ch := make(chan string) // Unbuffered channel
	go func() {
		ch <- "Hello from goroutine" // Send message
	}()

	// The main goroutine will block here until the message is received
	fmt.Println(<-ch)
	// foo:= <-ch

	/* buffered channels */
	// ch := make(chan int, 3) // Buffered channel with a capacity of 3

	// ch <- 1
	// ch <- 2
	// ch <- 3

	// // The channel is full now. Any further send operation will block.

	// fmt.Println(<-ch) // Receive 1
	// fmt.Println(<-ch) // Receive 2
	// fmt.Println(<-ch) // Receive 3

	/* select statement */
	// ch1 := make(chan string)
	// ch2 := make(chan string)

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	ch1 <- "Message from ch1"
	// }()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	ch2 <- "Message from ch2"
	// }()

	// select {
	// case msg1 := <-ch1:
	// 	fmt.Println("Received:", msg1)
	// case msg2 := <-ch2:
	// 	fmt.Println("Received:", msg2)
	// case <-time.After(3 * time.Second):
	// 	fmt.Println("Timeout")
	// }

	/* ranging over channels */
	// ch := make(chan int)

	// // Sender goroutine
	// go func() {
	// 	for i := 1; i <= 5; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch) // Close the channel when done sending
	// }()

	// // Receiver: Range over the channel
	// for val := range ch {
	// 	fmt.Println(val)
	// }

	// fmt.Println("Channel closed, all values received.")
}
