package main

import "fmt"

func main() {
	fmt.Println("For range over int")
	forRangeOverInt()
}

func forRangeOverInt() {
	for i := range 10 {
		fmt.Println(10 - i)
	}
}
