package main

import (
	"fmt"
	"os"
)

func main() {
	// Get an environment variable
	path := os.Getenv("PATH")
	fmt.Println("PATH:", path)

	// // Create a file
	// file, err := os.Create("example.txt")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// defer file.Close()

	// // Write to the file
	// file.WriteString("Hello, Go!")

	// // Read from the file
	// content, _ := os.ReadFile("example.txt")
	// fmt.Println("File Content:", string(content))

	// // Delete the file
	// os.Remove("example.txt")

}
