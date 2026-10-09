package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func calculate(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	a := 12
	b := 14

	// ---------------------

	// sum := add(a, b)
	// product := multiply(a, b)

	// fmt.Println("The sum is", sum)
	// fmt.Println("The product is", product)

	// ---------------------

	// sum, product := calculate(a, b)

	// fmt.Println("The sum is", sum)
	// fmt.Println("The product is", product)

	// ---------------------

	_, product := calculate(a, b)
	fmt.Println("The product is", product)

}
