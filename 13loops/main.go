package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops in GoLang")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i, _ := range days {
		fmt.Println(days[i])
	}

	for idx, day := range days {
		fmt.Printf("index is %v and %v\n", idx, day)
	}

	rougueValue := 1
	for rougueValue < 10 {
		fmt.Println(rougueValue)
		rougueValue++
		if rougueValue == 2 {
			goto label
		}
	}

label:
	fmt.Println("HEY HIMANK")

}
