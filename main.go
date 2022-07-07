package main

import (
	"github.com/tel21a-inf2/benchmarks/sorting"
)

const (
	iterations = 100
)

func main() {
	//branchprediction.RunBenchmark(200, 30000, iterations)
	sorting.RunBenchmark(200, 30000, iterations)
	// l1 := []int{5, 235, 2, 3, 42, 1, 107, 7, 12}
	// sorting.QuickSort(l1)
	// fmt.Println(l1)
}
