package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("LCO web Requests")

	response, err := http.Get(url)

	checkNilError(err)

	fmt.Printf("Response type: %T\n", response)

	defer response.Body.Close() //need to be close manually

	dataByte, err := ioutil.ReadAll(response.Body)

	checkNilError(err)

	fmt.Println("Response Body:", string(dataByte))

}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
