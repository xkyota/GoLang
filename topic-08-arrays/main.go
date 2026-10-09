package main

import "fmt"

func main() {
	// numbers := [...]int{10, 20, 30, 40, 50}

	// fmt.Println(numbers[1])

	// numbers[1] = 100
	// fmt.Println(numbers[1])

	// fmt.Println("Array length is", len(numbers))

	//! ====================

	// a := [3]int{10, 20, 30}
	// b := a

	// b[0] = 100

	// fmt.Println("A: ", a)
	// fmt.Println("B: ", b)

	//! ====================

	requests := [5]int{120, 85, 200, 150, 95}
	var numberOfRequests int
	maxRequests := requests[0]
	maxRequestsIndex := 0
	fmt.Println("Array length:", len(requests))

	for index, value := range requests {
		fmt.Println(index, value)
		numberOfRequests += value

		if value > maxRequests {
			maxRequests = value
			maxRequestsIndex = index
		}
	}

	fmt.Println("Number of requests:", numberOfRequests)
	fmt.Println("Maximum requests per day:", maxRequests)
	fmt.Println("Index of maximum requests:", maxRequestsIndex)

	//! ====================

	backup := requests
	backup[0] = 999

	fmt.Println("Requests:", requests)
	fmt.Println("Backup:", backup)

}
