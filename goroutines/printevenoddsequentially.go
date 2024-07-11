package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	evenChan, oddChan := make(chan bool, 1), make(chan bool, 1)
	wg.Add(2)
	go printOdd(evenChan, oddChan)
	go printEven(evenChan, oddChan)

	evenChan <- true
	wg.Wait()
	close(evenChan)
	close(oddChan)
}

func printOdd(evenChan, oddChan chan bool) {
	for i := 1; i < 11; i = i + 2 {
		<-oddChan
		fmt.Println(i)
		evenChan <- true
	}
	wg.Done()
}

func printEven(evenChan, oddChan chan bool) {
	for i := 0; i < 11; i = i + 2 {
		<-evenChan
		fmt.Println(i)
		oddChan <- true
	}
	wg.Done()
}
