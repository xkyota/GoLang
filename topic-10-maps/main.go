package main

import (
	"fmt"
)

func main() {
	// ages := map[string]int{
	// 	"Alex": 20,
	// 	"Max":  25,
	// 	"John": 30,
	// }

	// fmt.Println(ages)

	// fmt.Println()

	// ages["Maksym"] = 17
	// fmt.Println(ages)

	// fmt.Println()

	// fmt.Println("Maksym is", ages["Maksym"], "years old")

	// fmt.Println()

	// delete(ages, "Alex")
	// fmt.Println(ages)

	// fmt.Println()

	// age, ok := ages["Andris"]
	// fmt.Println("Age:", age)
	// fmt.Println("Exists:", ok)

	// if age, ok := ages["Andris"]; ok {
	// 	fmt.Println("User age:", age)
	// } else {
	// 	fmt.Println("User not found")
	// }

	//! ====================

	requests := map[string]int{
		"/api/users":    120,
		"/api/products": 85,
		"/api/orders":   50,
	}

	fmt.Println("Number of requests to /api/users:", requests["/api/users"])

	requests["/api/login"] = 30
	requests["/api/users"] += 10

	delete(requests, "/api/orders")

	fmt.Println(len(requests))

	if count, ok := requests["/api/admin"]; ok {
		fmt.Println("Requests:", count)
	} else {
		fmt.Println("Route not found")
	}

	sum := 0
	for request, numberOfRequests := range requests {
		fmt.Println(request, numberOfRequests)
		sum += numberOfRequests
	}

	fmt.Println("Total requests:", sum)
}
