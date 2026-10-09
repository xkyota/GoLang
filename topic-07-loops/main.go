package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Line №", i)
	// }

	//! ====================

	// count := 1
	// for count <= 5 {
	// 	fmt.Println("Count", count)
	// 	count++
	// }

	//! ====================

	// count := 1

	// for {
	// 	fmt.Println("Line Number: ", count)

	// 	if count == 5 {
	// 		break
	// 	}

	// 	count++
	// }

	//! ====================

	// numbers := []int{10, 20, 30, 40, 50}

	// for index, valude := range numbers {
	// 	fmt.Println(index, valude)
	// }

	//! ====================

	numbers := []int{4, 7, 12, 3, 18, 9, 6}
	var sum int
	var evens int

	for _, value := range numbers {
		sum += value

		if value%2 == 0 {
			evens++
		}
	}

	fmt.Println("Sum of all numbers:", sum)
	fmt.Println("Number of even numbers:", evens)

	//! ====================

	for number := 0; number < 10; number++ {
		if number == 5 {
			continue
		} else if number == 8 {
			break
		}

		fmt.Println(number)
	}

}
