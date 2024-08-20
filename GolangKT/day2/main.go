package main

import "os"

func main() {
	file, err := os.Open("file.txt")
	if err != nil {
		//
	}
	defer file.Close()

}

func logRequestTime() {

	// current := time.Time()

	// defer time.Time() - current

	// process request

}

// func Add() {
// 	var n int
// }
