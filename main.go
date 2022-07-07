package main

import (
	"github.com/tel21a-inf2/benchmarks/branchprediction"
	"github.com/tel21a-inf2/benchmarks/lists"
	"github.com/tel21a-inf2/benchmarks/sorting"
)

const (
	maxElement = 200
	listLength = 30000
	iterations = 100
)

func main() {
	branchprediction.RunBenchmark(maxElement, listLength, iterations)
	sorting.RunBenchmark(maxElement, listLength, iterations)
	lists.RunBenchmark(maxElement, listLength, 5, iterations)
}
