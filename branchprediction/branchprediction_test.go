package branchprediction

import "fmt"

func ExampleSumGreater() {
	l1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i <= 10; i++ {
		fmt.Println(SumGreater(l1, i))
	}

	// Output:
	// 55
	// 54
	// 52
	// 49
	// 45
	// 40
	// 34
	// 27
	// 19
	// 10
	// 0
}
