package main

import (
	"chapter_02/tempconv"
	"fmt"
)

// type Celsius float64
// type Farhanheit float64

// func main() {
// 	var i, j, k = 10, true, "abc"

// 	fmt.Println(i, j, k)

// 	a, b := 1, 0

// 	fmt.Println(a, b)

// 	a, b = b, a

// 	fmt.Println(a, b)

// }

// func CToF(c Celsius) Farhanheit {
// 	return Farhanheit(c*9/5 + 32)
// }

// func FToC(f Farhanheit) Celsius {
// 	return Celsius((f - 32) * 5 / 9)
// }

func test() string {
	return "x"
}

func main() {

	convertedToF := tempconv.FToC(100.00)
	fmt.Println(convertedToF)
	// fmt.Println(reflect.TypeOf(convertedToF))
	// // fmt.Printf("%g\n", convertedToF)
	// var convertedToC Farhanheit = 0

	// // fmt.Println(convertedToC, convertedToF)
	// // fmt.Printf("%g", convertedToC)

	// fmt.Println(convertedToC == Farhanheit(convertedToF))

	// a := "123"
	// p := &a
	// fmt.Println(p)
	// fmt.Println(*p)
	// x := test
	// fmt.Println(reflect.TypeOf(&))

	// var p = funn()

	// fmt.Println(funn())
	// fmt.Println(funn())
	// fmt.Println(funn())
	// fmt.Println(funn() == funn())
	// v := 10
	// incr(&v)
	// fmt.Println(v)
	// fmt.Println(incr(&v))

	// fmt.Println(incr(&v))

	// p := new(100)
	// q := new(999)
	// fmt.Println(p == q)

	// a, b, _ := "a", 12, ""

	// fmt.Println(a, b)
}
func new(number int) int {
	return number * number
}
func funn() *int {
	v := 111
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
