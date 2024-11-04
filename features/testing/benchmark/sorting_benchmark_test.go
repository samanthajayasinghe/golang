package main

import (
	"math/rand"
	"testing"
)

func makeRandomNumberSlice(n int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(n)
	}
	return numbers
}

// const LENGTH = 10_000
const LENGTH = 100

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		numbers := makeRandomNumberSlice(LENGTH)

		b.StartTimer()
		BubbleSort(numbers)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		numbers := makeRandomNumberSlice(LENGTH)

		b.StartTimer()
		QuickSort(numbers)
	}
}
