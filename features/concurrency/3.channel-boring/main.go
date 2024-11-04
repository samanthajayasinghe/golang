package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string)

	go boring("boring A", c)
	go boring("boring B", c)

	for i := 0; i < 10; i++ {
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

/////var c chan int

// sending to channel
//c <- 1

// Receiving from channel
//value := <-c
