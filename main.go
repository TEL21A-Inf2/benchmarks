package main

import "github.com/tel21a-inf2/benchmarks/branchprediction"

func main() {
	branchprediction.RunBenchmark(100, 100)
	branchprediction.RunBenchmark(100, 1000)
	branchprediction.RunBenchmark(100, 10000)
	branchprediction.RunBenchmark(100, 100000)
	branchprediction.RunBenchmark(100, 1000000)
}
