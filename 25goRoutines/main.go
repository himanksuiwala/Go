package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var websiteList = []string{
	"https://github.com",
	"https://fb.com",
	"https://amazon.com",
	"https://google.com",
}
var wg sync.WaitGroup
var signals = []string{"test"} //pointer
var mut sync.Mutex             //pointer

func main() {
	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)

	// fmt.Println("Go Routine Tutorial")
	// go greeter("hello")
	// greeter("World")
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPs in Endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}

func greeter(s string) {
	time.Sleep(0)
	for i := 0; i < 6; i++ {
		fmt.Println(s)
	}
}
