package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	result := FetchUsers()

	// go object to json

	// response, err := json.Marshal(result)
	// if err != nil {
	// 	fmt.Fprint(w, "Some error occured")
	// }

	w.Header().Set("content-type", "application/json")
	// fmt.Fprint(w, string(response))

	json.NewEncoder(w).Encode(result)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user Users

	// io.Reader
	// io.Writer
	json.NewDecoder(r.Body).Decode(&user)

	DB.Create(user)
}
