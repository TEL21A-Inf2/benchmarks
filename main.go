package main

import "github.com/tel21a-inf2/benchmarks/branchprediction"

const (
	iterations = 1000
)

func main() {
	branchprediction.RunBenchmark(200, 10000, iterations)
}
