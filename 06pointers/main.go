package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on Pointers")

	// var ptr *int
	// fmt.Println("Default value of Ptr:", ptr)

	myNumber := 24
	var ptr = &myNumber
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = *ptr + 2
	fmt.Println("NEW VALUE", myNumber)
}
