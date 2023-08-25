package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() bool {
	return u.Status
}
func (u User) newMail() {
	u.Email = "t@d.com"
	fmt.Println(u.Email)
}

func main() {
	fmt.Println("Structs of Golang")

	himank := User{"Himank", "himank@dev.com", true, 16}
	fmt.Println(himank)
	fmt.Printf("Details: %+v\n", himank)
	userStatus := himank.getStatus()
	fmt.Println(himank.Email)
	himank.newMail()
	fmt.Println(userStatus, himank.Email)
}
