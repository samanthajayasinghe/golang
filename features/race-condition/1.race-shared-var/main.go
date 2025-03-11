package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup

func increment() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		counter++
	}
}

func main() {
	wg.Add(2)
	go increment()
	go increment()
	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
