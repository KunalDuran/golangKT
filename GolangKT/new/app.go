package main

import "fmt"

func main() {
	fmt.Println("----------shallow copy")
	products := make([]int, 1)
	products[0] = 1
	products = append(products, 2)
	products = append(products, 3)
	fmt.Printf("%p\n", products)

	newProducts := copy(products)
	products = append(products, 4)
	newProducts[0] = 99
	fmt.Println("products are ", products)
	fmt.Println("new products are ", newProducts)

	fmt.Printf("%p\n", products)
	fmt.Printf("%p\n", newProducts)
}

func copy(products []int) []int {
	return products
}
