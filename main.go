package main

import (
	"github.com/tel21a-inf2/benchmarks/branchprediction"
	"github.com/tel21a-inf2/benchmarks/lists"
	"github.com/tel21a-inf2/benchmarks/sorting"
)

const (
	iterations = 100
)

func main() {
	branchprediction.RunBenchmark(200, 30000, iterations)
	sorting.RunBenchmark(200, 30000, iterations)
	lists.RunBenchmark(200, 30000, 5, iterations)
}
