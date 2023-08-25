package main

import "fmt"

func main() {
	fmt.Println("Arrays in GO")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Apple"
	fruitList[2] = "Banana"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List:", fruitList)

	var vegList = [3]string{"Potato", "TP,atp", "PAlak"}
	fmt.Println(vegList)
}
