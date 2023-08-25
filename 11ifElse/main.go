package main

import "fmt"

func main() {
	fmt.Println("If else in Golang")

	loginCount := 23
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Not a regular User"
	} else {
		result = "Not a not regular Message"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("EVEN")
	} else {
		fmt.Println("ODD")

	}

	if num := 3; num < 10 {
		fmt.Println("Less than 10")
	}

	err := ""
	if err != "OK" {
		fmt.Println("as")

	}
}
