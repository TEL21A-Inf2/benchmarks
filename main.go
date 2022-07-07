package main

import "github.com/tel21a-inf2/benchmarks/branchprediction"

func main() {
	branchprediction.RunBenchmark(1, 10000, 5000)
	branchprediction.RunBenchmark(1, 1000000, 500000)
	branchprediction.RunBenchmark(1, 100000000, 50000000)
}
