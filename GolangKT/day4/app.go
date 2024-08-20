package main

func main() {
	// channels use-case with example

	// waiters - take orders from customers
	// chefs - prepare the order

	// problem statement - communicate order details via channel

	orderChan := make(chan Order)

	for i := 0; i < 3; i++ { // represents 3 waiters
		takeOrder(orderChan)
	}

	// logic to prepare order (chef)

}

type Order struct {
	Details string
	Table   int
}

func takeOrder(ch chan Order) {
	//
}

func prepareOrder(ch chan Order) {
	//
}
