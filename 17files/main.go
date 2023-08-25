package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in golang ")
	content := "To be inside File"

	file, err := os.Create("./filemadeusingGO.txt")

	checkNilError(err)
	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("Length:", length)
	defer file.Close()
	readFile(file.Name())
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)

	checkNilError(err)
	fmt.Println("Text inside the File \n", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
