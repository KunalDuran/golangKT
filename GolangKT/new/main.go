package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// 	// foo := 1
// 	// fmt.Printf("%T\n", foo)
// 	// fmt.Printf("%T", MyInt(foo))

// 	workCh := make(chan string)
// 	workCh2 := make(chan string)

// 	go Worker(workCh)
// 	go Worker2(workCh2)

// 	for {
// 		select {
// 		case msg := <-workCh:
// 			fmt.Println("msg from worker " + msg)
// 		case msg := <-workCh2:
// 			fmt.Println("msg from worker " + msg)
// 		}
// 	}
// }

// type MyInt int

// func Worker(ch chan string) {
// 	for {
// 		ch <- "1"
// 		time.Sleep(time.Second)
// 	}
// }

// func Worker2(ch chan string) {
// 	for {
// 		ch <- "2"
// 		time.Sleep(time.Second * 2)
// 	}
// }
