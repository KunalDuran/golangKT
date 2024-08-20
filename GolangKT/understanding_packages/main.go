package main

import "net/http"

func init() {
	InitDB()
}

func main() {
	// - handling json
	// - set headers
	// - conn to db
	// - routes declaration
	http.HandleFunc("/", Home)
	http.HandleFunc("/users", HandleUsers)
	http.HandleFunc("/create-user", CreateUser)

	http.ListenAndServe(":8080", nil)
}
