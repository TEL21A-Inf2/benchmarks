package main

import "github.com/tel21a-inf2/benchmarks/branchprediction"

func main() {
	branchprediction.RunBenchmark(100, 100, 50)
	branchprediction.RunBenchmark(100, 1000, 500)
	branchprediction.RunBenchmark(100, 10000, 5000)
	branchprediction.RunBenchmark(100, 100000, 50000)
	branchprediction.RunBenchmark(100, 1000000, 500000)
}
