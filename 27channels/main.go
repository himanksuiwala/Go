package main

import (
	"fmt"
	"sync"
)

func main() {
	// fmt.Println("Channel in GO")

	myChannel := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func(myChannel <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChannel

		fmt.Println(isChannelOpen)
		fmt.Println(val)
		fmt.Println(<-myChannel)
		fmt.Println(<-myChannel)
		wg.Done()
	}(myChannel, wg)
	go func(myChannel chan<- int, wg *sync.WaitGroup) {
		myChannel <- 5
		myChannel <- 6

		wg.Done()
	}(myChannel, wg)

	wg.Wait()
}
