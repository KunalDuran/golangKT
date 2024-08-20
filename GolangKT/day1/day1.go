package main

import (
	"fmt"
	"math"
)

func main() {
	age := 31
	_ = age
	/* */
	//

	// zero value

	// -Boolean types
	var paymentDone bool // true, false
	paymentDone = true
	fmt.Println(paymentDone)
	fmt.Println(&paymentDone)

	// -Numeric types
	var Price float32
	fmt.Println(Price)

	// var CartItems int
	// short hand declaration
	CartItems := 5
	fmt.Println(CartItems)

	// -String types
	// var productName string
	productName := "Bat"
	fmt.Printf("product name : %s \n", productName)

	// -Array types
	// var product [2]string
	// product[0] = "Bat"
	// product[1] = "Ball"

	product := [2]string{"Bat", "Ball"}
	fmt.Println(product)

	// -Slice types
	// prices := []string{"Bat", "Ball"}

	// built-in to add items to the slice
	// prices = append(prices, "Wicket")
	// fmt.Println(prices)

	prices := make([]int, 2, 5)
	prices[0] = 3
	fmt.Println(prices)

	// fmt.Printf("%p \n", &prices)

	// prices = append(prices, 3)

	// fmt.Printf("%p \n", &prices)

	// prices = append(prices, 7)

	// fmt.Printf("%p \n", &prices)

	// prices = append(prices, 10)

	// fmt.Printf("%p \n", &prices)

	// // prices = append(prices, 10)
	// fmt.Println(prices)

	// -Pointer types

	// -Struct types
	// var prod Product
	// prod.Name = "Bat"
	// prod.Price = 19

	prod := Product{
		Name:  "Bat",
		Price: 4,
	}

	fmt.Println(prod)

	// -Map types
	var cart = make(map[string]int)
	cart["Bat"] = 100
	fmt.Println(cart)

	// -Function types
	result, err := Add(1, 4)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)

	// methods

	var sq Square

	sq.Height = 4

	fmt.Println(sq.Area())

	var c Circle

	c.Radius = 2
	fmt.Println(c.Area())

	var r Rectange
	r.Length = 3
	r.Breadth = 2

	fmt.Println(r)

	// -Interface types
	CalculateArea(sq)
}

type Circle struct {
	Radius int
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.Radius) * float64(c.Radius)
}

type Square struct {
	Height float64
}

func (s Square) Area() float64 {
	return s.Height * s.Height
}

type Rectange struct {
	Length  int
	Breadth int
}

type Shape interface {
	Area() float64
}

func Add(a, b int) (int, error) {
	return a + b, nil
}

type Product struct {
	Name    string
	Price   int
	Details Information
}

type Information struct {
	Message string
}

func CalculateArea(s Shape) float64 {
	return s.Area()
}
