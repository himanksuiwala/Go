package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const myurl = "https://google.com"

type course struct {
	Name     string
	Price    int
	platform string
	Password string
	Tags     []string
}

func EncodeJson() {
	lcoCourse := []course{
		{"React Bootcamp", 299, "LCO.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootamp", 199, "LCO.in", "abc123", []string{"web-dev", "Full-stack", "js"}},
		{"MongoDB", 129, "LCO.in", "abc123", nil},
	}

	finalJson, err := json.Marshal(lcoCourse)

	checkNilError(err)
	fmt.Print("%s\n", finalJson)

}

func main() {
	fmt.Println("Make web requests- CRUD using Python")
	// getRequest()
	// performFormRequest()
	EncodeJson()
}

func getRequest() {

	response, err := http.Get(myurl)
	checkNilError(err)
	// fmt.Println("Response", response)
	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Status Code:", response.ContentLength)

	// content, err := ioutil.ReadAll(response.Body)

	// fmt.Println(string(content))
}

func postRequest() {
	reqBody := strings.NewReader(`
		{
			"coursename:"Course on GO",
			"price": 100,
			"platform","learncode.com
		}
	`)

	response, err := http.Post(myurl, "application/json", reqBody)
	checkNilError(err)
	defer response.Body.Close()
	ioutil.ReadAll(response.Body)
}

func performFormRequest() {
	data := url.Values{}
	data.Add("firstname", "Him")
	data.Add("lastname", "Coco")
	data.Add("mail", "mail.com")

	response, err := http.PostForm(myurl, data)

	checkNilError(err)
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
