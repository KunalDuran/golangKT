package main

type User struct {
	FirstName string `json:"first_name" gorm:"first_name"`
	LastName  string `json:"last_name" gorm:"last_name"`
	Email     string `json:"email" gorm:"email"`
	Password  string `json:"password" gorm:"password"`
}
