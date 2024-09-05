package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)

	go boring("boring", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)

	}
	fmt.Print("You are boring; I'm leaving")

}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(1 * time.Millisecond))
	}
}

// ========== Notes ==============

//var c chan string

// sending to channel
//c <- 1

// Receiving from channel
//value := <-c
