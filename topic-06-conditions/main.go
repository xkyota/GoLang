package main

import "fmt"

func main() {

	port := 8080
	isProduction := false
	command := "start"

	if isProduction == true && port == 8080 {
		fmt.Println("Production server")
	} else if isProduction == false && port == 8080 {
		fmt.Println("Development server")
	} else {
		fmt.Println("Custom configuration")
	}

	switch command {
	case "start":
		fmt.Println("Starting Server...")
	case "stop":
		fmt.Println("Stopping Server...")
	case "restart":
		fmt.Println("Restarting Server...")
	}

}
