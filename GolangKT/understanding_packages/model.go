package main

type Users struct {
	FirstName string `json:"first_name,omitempty" xml:""`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email"`
	Password  string `json:"-"`
}
