package main

import (
	"fmt"
)

func main() {
	printExample("Min of single integer")
	minOfSingleInt()
	printExample("Min of multiple integers")
	minOfMultipleInt()
	printExample("Min of multiple strings")
	minOfMultipleString()
}

func minOfSingleInt() {
	x := 1
	m := min(x)
	fmt.Printf("min of %d is %d \n", x, m)
}

func minOfMultipleInt() {
	var x, y, z int
	x = -1
	y = 0
	z = 1
	m := min(x, y, z)
	fmt.Printf("min of %d,%d,%d is %d \n", x, y, z, m)
}

func minOfMultipleString() {
	var a, b, c string
	a = "foo"
	b = "bar"
	c = "abc"
	m := min(a, b, c)
	fmt.Printf("min of %s,%s,%s is %s \n", a, b, c, m)
}

func printExample(header string) {
	fmt.Println("\n=======================================")
	fmt.Println(header)
	fmt.Println("=======================================")
}
