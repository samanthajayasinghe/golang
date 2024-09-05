package main

import (
	"fmt"
	"time"
)

func main() {

	//simple()
	handleAsService()
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() { // launch goroutine inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(1 * time.Millisecond))
		}
	}()
	return c // return the channel
}

func simple() {
	c := boring("boring")

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)

	}
	fmt.Print("You are boring; I'm leaving")
}

func handleAsService() {
	tim := boring("Tim")
	ann := boring("Ann")

	for i := 0; i < 5; i++ {
		fmt.Println(<-tim)
		fmt.Println(<-ann)
	}
	fmt.Print("You are boring; I'm leaving")
}
