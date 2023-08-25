package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to video on slice")

	var fruitList = []string{"Tomato", "Apple", "Peach"}
	fmt.Printf("Type of FruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango")
	fmt.Println("Type of FruitList is", fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScore := make([]int, 4)

	highScore[0] = 234
	highScore[1] = 945
	highScore[2] = 465
	highScore[3] = 867
	// highScore = append(highScore, 1, 2, 3, 4, 5, 6, 7) /

	// fmt.Println(sort.IntsAreSorted(highScore))
	// sort.Ints(highScores)
	// sort.Ints(highScore)
	// fmt.Printf("%v", highScore)

	//To remove a value based upon the index
	courses := []string{"JS", "React", "GO", "Node.js", "Angular"}
	fmt.Println(courses)
	courses = append(courses[:2], courses[2+1:]...)

	fmt.Printf("%T", courses[2+1:])
}
