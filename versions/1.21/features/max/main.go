package main

import (
	"fmt"
)

func main() {
	printExample("Max of single integer")
	maxOfSingleInt()
	printExample("Max of multiple integers")
	maxOfMultipleInt()
	printExample("Max of multiple strings")
	maxOfMultipleString()
}

func maxOfSingleInt() {
	x := 1
	m := max(x)
	fmt.Printf("max of %d is %d \n", x, m)
}

func maxOfMultipleInt() {
	var x, y, z int
	x = -1
	y = 0
	z = 1
	m := max(x, y, z)
	fmt.Printf("max of %d,%d,%d is %d \n", x, y, z, m)
}

func maxOfMultipleString() {
	var a, b, c string
	a = "foo"
	b = "bar"
	c = "abc"
	m := max(a, b, c)
	fmt.Printf("max of %s,%s,%s is %s \n", a, b, c, m)
}

func printExample(header string) {
	fmt.Println("\n=======================================")
	fmt.Println(header)
	fmt.Println("=======================================")
}
