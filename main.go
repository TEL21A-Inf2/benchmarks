package main

import "github.com/tel21a-inf2/benchmarks/branchprediction"

const (
	iterations = 100
)

func main() {
	branchprediction.RunBenchmark(200, 30000, iterations)
}
