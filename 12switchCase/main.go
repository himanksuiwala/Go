package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Switch-Case in GoLang")
	// New(NewSource(seed))
	diceNumber := rand.Intn(6) + 1
	fmt.Println(diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 & you can open")
		fallthrough
	case 2:
		fmt.Println("Dice Value is 2 & you cannot open")
		fallthrough
	case 3:
		fmt.Println("Dice Value is 3 & you cannot open")
	case 4:
		fmt.Println("Dice Value is 4 & you cannot open")
	case 5:
		fmt.Println("Dice Value is 5 & you cannot open")
	case 6:
		fmt.Println("Dice Value is 6 & you can open")
	default:
		fmt.Println("Default Value")

	}
}
