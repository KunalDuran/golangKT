package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {

// 	ch := make(chan int)

// 	go func() {
// 		for i := 0; i < 50; i++ {
// 			ch <- WebRequest("http://google.com")
// 		}
// 		close(ch)
// 	}()

// 	// range over channel
// 	for msg := range ch {
// 		fmt.Println(msg)
// 	}

// 	// for {
// 	// 	foo, ok := <-ch
// 	// 	if !ok {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(foo)
// 	// }
// }

// func WebRequest(url string) int {
// 	// request to the url
// 	// send status code to channel
// 	return rand.Int()
// }

// // single core
// // 1,2,4,68,5,6,7,8,9,90

// // get request
// // 3
