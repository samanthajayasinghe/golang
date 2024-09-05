package main

import (
	"fmt"
	"time"
)

func main() {

	go boring("boring")
	//time.Sleep(time.Duration(5 * time.Millisecond))
	//fmt.Print("You are boring; I'm leaving")
}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(1 * time.Millisecond))
	}
}
