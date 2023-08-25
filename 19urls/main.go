package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&payment=paytm"

func main() {
	fmt.Println("Handling URLs in golang")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println("Params of query", qparams)

	for _, idx := range qparams {
		fmt.Println("Params:", idx)
	}
	partsofUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}

	anotherUrl := partsofUrl.String()

	fmt.Println(anotherUrl)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
