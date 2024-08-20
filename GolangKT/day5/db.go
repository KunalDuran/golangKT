package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() {
	dsn := "root:password@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

}

func fetchUser() {
	var user User

	if err := DB.Find(&user).Error; err != nil {
		panic(err)
	}

	fmt.Println(user)
}

func createUser() {

	user := User{
		FirstName: "Kunal",
		LastName:  "Kumar",
		Email:     "kunal@kt.com",
		Password:  "password",
	}

	DB.Create(&user)

	fmt.Println(user)
}
