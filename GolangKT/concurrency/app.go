package main

import (
	"fmt"
	// "sync"
	"time"
)

func main() {

	// var wg sync.WaitGroup

	// wg.Add(2)

	// go func() {
	// 	defer wg.Done()
	// 	StreamingService("video")
	// }()

	// go func() {
	// 	defer wg.Done()
	// 	go StreamingService("audio")
	// }()

	// wg.Wait()

	// ch := make(chan string, 5) // buffered
	// ch := make(chan string) // unbuffered
	// go func() {
	// 	fmt.Println(<-ch)
	// }()
	// ch <- "Hello"

	// // time.Sleep(time.Second)
	// fmt.Println("Exiting program")

	// ch := make(chan int, 3) // Buffered channel with a capacity of 3

	// ch <- 1
	// ch <- 2
	// ch <- 3

	// // The channel is full now. Any further send operation will block.

	// fmt.Println(<-ch) // Receive 1
	// fmt.Println(<-ch) // Receive 2
	// fmt.Println(<-ch) // Receive 3

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

func WebRequest(url string, ch chan int) {
	// request to the url
	// send status code to channel
	var statusCode int
	ch <- statusCode
}

func StreamingService(name string) {
	for {
		fmt.Println("streaming " + name)
		time.Sleep(time.Millisecond * 500)
	}
}
