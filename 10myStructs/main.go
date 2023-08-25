package main

import (
	"fmt"
)

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs of Golang")

	himank := User{"Himank", "himank@dev.com", true, 16}
	fmt.Println(himank)
	fmt.Printf("Details: %+v\n", himank)
}
