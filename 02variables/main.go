package main

import "fmt"

var LoginToken string = "123123123" //Capital L makes this a public variable

func main() {
	var username string = "Himank"
	var isLoggedIn bool = true
	fmt.Println(username)

	fmt.Printf("Variable is of type: %T \n", isLoggedIn)
	fmt.Printf("Variable is of type: %T, is Logged In: %T \n", username, isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println("Uint8:", smallVal)

	var smallFloat float32 = 255.12123132132123
	fmt.Println("Uint8:", smallFloat)

	//default values and some aliases
	var anotherVariable string
	fmt.Println(anotherVariable)

	var website = "google.com"
	fmt.Printf("Website is of type:'%s'", website)

	numberofUser := 1234566
	fmt.Printf("number of Users : '%c'", numberofUser)

	fmt.Println(LoginToken)
}
