package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

const (
	a = 200
	d
	b = 300
	e
)

type Weekday rune

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Employee struct {
	ID            int
	Name, Address string
	Position      string
	Salary        int
}

type Point struct {
	X, Y int
}

type address struct {
	hostname string
	port     int
}

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty,"`
	Actors []string
}

func bonus(employee *Employee) int {
	fmt.Printf("%p \n", &employee)
	return employee.Salary * 10
}
func main() {

	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
		// ...
	}
	data, _ := json.MarshalIndent(movies, "", "   ")
	fmt.Printf("%s", data)

	// hits := make(map[address]int)

	// hits[address{"google.com", 1001}] += 999

	// fmt.Println(hits[address{"google.com", 1001}])

	// p := Point{100, 1200}
	// q := Point{100, 200}

	// fmt.Println(p.Y == q.Y)

	// himank := Employee{12, "GoGo Watson", "Zurich", "Senior Go Advocate", 100}
	// fmt.Printf("%p \n", &himank)
	// fmt.Println(bonus(&himank))

	// var goMan Employee
	// goMan.ID = 123
	// goMan.Name = "Go-MAN !"
	// goMan.Position = "SuperMan"
	// goMan.Salary = 90000
	// goMan.Address = "New Delhi"

	// var moMan Employee

	// moMan = goMan
	// fmt.Printf("%p \n", &goMan)
	// fmt.Printf("%p", &moMan)
	// abc := map[string]int{"a": 123}
	// efg := map[string]int{"a": 123}
	// var empty = make(map[string]int)
	// fmt.Println(empty == nil)
	// // fmt.Println(abc == efg)
	// ages := map[string]int{
	// 	"alice": 19,
	// 	// "alice": 19,
	// 	"Bob": 22,
	// }
	// _, ok := ages["Boab"]
	// if !ok {
	// 	fmt.Println("NOT PRESENT")
	// }
	// delete(ages, "alice")
	// fmt.Println(ages["alice"], ages["Bob"])
	// result := []rune{}

	// for _, val := range "HELLO"{
	// 	result = append(result, val)
	// }

	// fmt.Printf("%q", result)
	// slice1 := [...]int{10, 20, 30, 40}
	// slice2 := [...]int{10, 20, 30, 40}

	// fmt.Println(reflect.TypeOf(slice1))
	// fmt.Println(slice1 == slice2)

	// var a [3]int
	// fmt.Println(a, b, d, e)

	// fmt.Println(Sunday)
	// fmt.Println(Monday)
	// fmt.Println(Tuesday)

	// for idx, val := range a {
	// 	fmt.Println(idx, val, "\n")
	// }

	// var q [3]int = [...]int{10, 20, 30}
	// 	q[] :=[...]int{1,2,4}
	// for _, val := range q {
	// 	fmt.Println(val, "\n")
	// }

	// q := [...]int{1, 2, 0, 123, 456, 678, 909, 1.}
	// // p := [3]int{1, 2}
	// q[0] = 101

	// fmt.Println(q[0])
	// // q = p

	// // a := [...]string{123: "abc", 456: "def", 789: "ghi"}

	// for idx, val := range q {
	// 	fmt.Println(idx, val)
	// }

	// fmt.Println(q[1:2])
	// // fmt.Println(p == q)

	// fmt.Println(intToString([]int{1, 2, 3, 4}))

}

func intToString(arr []int) string {
	var buffer bytes.Buffer
	buffer.WriteString("[")
	for idx, val := range arr {
		if idx > 0 {
			buffer.WriteString(", ")
		}
		fmt.Fprintf(&buffer, "%d", val)
	}
	buffer.WriteString("]")
	return buffer.String()
}

// a.go  a/b/c.go
func baseName(s string) string {

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1 : 0]
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}
