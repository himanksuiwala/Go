package main

import "fmt"

func main() {
	defer fmt.Println("first")
	defer fmt.Println("second")
	defer fmt.Println("Third")

	fmt.Println("Normal Statement")
	defer myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
