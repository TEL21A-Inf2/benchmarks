package main

import "github.com/tel21a-inf2/benchmarks/branchprediction"

const (
	iterations = 10
)

func main() {
	branchprediction.RunBenchmark(100, 100, iterations)
	branchprediction.RunBenchmark(100, 1000, iterations)
	branchprediction.RunBenchmark(100, 10000, iterations)
	branchprediction.RunBenchmark(100, 100000, iterations)
	branchprediction.RunBenchmark(100, 10000000, iterations)
}
