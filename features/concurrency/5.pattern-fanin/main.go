package main

import (
	"fmt"
	"time"

	"golang.org/x/exp/rand"
)

func main() {

	c := fanIn(boring("Tim"), boring("Ann"))

	for i := 0; i < 10; i++ {
		fmt.Printf("You say: %q\n", <-c)

	}
	fmt.Print("You are boring; I'm leaving")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() { // launch goroutine inside the function
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c // return the channel
}

// Note
// Multiplexing
