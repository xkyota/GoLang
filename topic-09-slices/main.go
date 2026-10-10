package main

import "fmt"

func main() {
	// numbers := []int{10, 20, 30}
	// fruits := make([]int, 3)

	// for index, value := range numbers {
	// 	fmt.Println(index, value)
	// }

	// fmt.Println()

	// for index, value := range fruits {
	// 	fmt.Println(index, value)
	// }

	//! ==================== Capacity

	// numbers := make([]int, 3, 5)
	// fmt.Println("Length:", len(numbers))
	// fmt.Println("Capacity:", cap(numbers))

	//! ==================== Append

	// numbers := []int{10, 20, 30}
	// numbers = append(numbers, 40)
	// fmt.Println(numbers)

	// numbers = append(numbers, 50, 50, 70)
	// fmt.Println(numbers)

	//! ==================== Slicing

	// numbers := []int{10, 20, 30, 40, 50}

	// partOfNumbers := numbers[1:3]
	// fmt.Println(partOfNumbers)

	//! ====================

	requests := []int{120, 85, 200}
	fmt.Println(requests)

	fmt.Println("Slice length:", len(requests))
	fmt.Println("Slice capacity:", cap(requests))

	fmt.Println()

	requests = append(requests, 150, 95)
	fmt.Println(requests)
	fmt.Println("Slice length:", len(requests))
	fmt.Println("Slice capacity:", cap(requests))

	fmt.Println()

	var sum int
	for _, value := range requests {
		sum += value
	}
	fmt.Println("Sum of all requests:", sum)

	fmt.Println()

	selected := requests[1:4]
	fmt.Println(selected)
	selected[0] = 999

	fmt.Println(requests)
	fmt.Println(selected)
	// Selected is using the same memory requests using. So element has been changed in both slices

	fmt.Println()

	backup := make([]int, len(requests))
	copy(backup, requests)
	backup[0] = 500

	fmt.Println(requests)
	fmt.Println(backup)

}
