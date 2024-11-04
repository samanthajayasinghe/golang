# Benchmark Test 

## How to run benchmark tests

```
go test -bench=.

go test -bench=. -cpu=1,2,4,12 
```

## How to read results 

The suffix, -16, in the first column (BenchmarkBubbleSort-16/BenchmarkQuickSort-16), tells us how many CPUs is used to run the tests. The second column is how many loops that was used during the test, b.N. The last column gives us the speed for each loop.

From this test we can clearly see who is the winner, QuickSort. Quicksort has an average-case time complexity of O(n log n), while BubbleSort has a worst-case and average-case time complexity of O(n^2).
