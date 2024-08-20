package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	WebRequest()
}

func WebRequest() {
	resp, err := http.Get("https://www.jsr-intlfzc.com/video/pc.mp4")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// body, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(body))

	file, err := os.Create("video.mp4")
	if err != nil {
		fmt.Println(err)
	}

	// b, err := file.Write(body)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("written %d bytes to file", b)

	buffer := make([]byte, 1024)

	for {
		b, err := resp.Body.Read(buffer)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("read %d bytes to file \n", b)

		// write
		file.Write(buffer)
	}
}
