package main

import (
	"fmt"
)

func main() {
	printExample("clear slice")
	clearSlice()
	printExample("clear Map")
	clearMap()

}

func clearSlice() {
	primes := [...]int{0, 1, 2, 3, 5, 7, 11, 13}
	fmt.Printf("Before clearing slice %d \n", primes)
	clear(primes[0:2])
	fmt.Printf("After clearing 2:4 element of the slice %d \n", primes)
}

func clearMap() {
	commits := map[string]int{
		"init repo":       10001,
		"my first commit": 10002,
		"my last commit":  10001,
		"before sleep":    100000,
	}
	fmt.Println("Before clearing the map")
	for k, v := range commits {
		fmt.Println(k, "hash is", v)
	}

	clear(commits)
	fmt.Println("After clearing the map")
	for k, v := range commits {
		fmt.Println(k, "hash is", v)
	}
}

func printExample(header string) {
	fmt.Println("\n=======================================")
	fmt.Println(header)
	fmt.Println("=======================================")
}
