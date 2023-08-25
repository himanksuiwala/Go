package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greeter()

	result := proAdder(1, 2, 3, 54, 68, 8, 8, 8, 8)
	result, msg := adder(1, 2)

	// result := adder(3, 5)

	fmt.Println("Result:", result, msg)

}
func adder(val1 int, val2 int) (int, string) {
	return val1 + val2, "HEYA"
}

func proAdder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func greeterTwo() {
	fmt.Println("Another method")

}
func greeter() {
	fmt.Println("Namastey from Golang")
}
