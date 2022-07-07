package branchprediction

import "fmt"

func ExampleSumGreater() {
	l1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(SumGreater(l1, 0))
	fmt.Println(SumGreater(l1, 1))
	fmt.Println(SumGreater(l1, 2))
	fmt.Println(SumGreater(l1, 3))
	fmt.Println(SumGreater(l1, 4))
	fmt.Println(SumGreater(l1, 5))
	fmt.Println(SumGreater(l1, 6))
	fmt.Println(SumGreater(l1, 7))
	fmt.Println(SumGreater(l1, 8))
	fmt.Println(SumGreater(l1, 9))
	fmt.Println(SumGreater(l1, 10))

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

func ExampleMultiples() {
	fmt.Println(Multiples(1, 5))
	fmt.Println(Multiples(2, 5))
	fmt.Println(Multiples(3, 5))
	fmt.Println(Multiples(10, 5))
	fmt.Println(Multiples(23, 5))

	fmt.Println(Multiples(1, 10))

	// Output:
	// [1 2 3 4 5]
	// [2 4 6 8 10]
	// [3 6 9 12 15]
	// [10 20 30 40 50]
	// [23 46 69 92 115]
	// [1 2 3 4 5 6 7 8 9 10]
}

func ExampleCreateListAndSum() {
	fmt.Println(CreateListAndSum(1, 10, 0))
	fmt.Println(CreateListAndSum(1, 10, 1))
	fmt.Println(CreateListAndSum(1, 10, 2))
	fmt.Println(CreateListAndSum(1, 10, 3))
	fmt.Println(CreateListAndSum(1, 10, 4))
	fmt.Println(CreateListAndSum(1, 10, 5))
	fmt.Println(CreateListAndSum(1, 10, 6))
	fmt.Println(CreateListAndSum(1, 10, 7))
	fmt.Println(CreateListAndSum(1, 10, 8))
	fmt.Println(CreateListAndSum(1, 10, 9))
	fmt.Println(CreateListAndSum(1, 10, 10))

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

func ExampleCreateListShuffleAndSum() {
	fmt.Println(CreateListShuffleAndSum(1, 10, 0))
	fmt.Println(CreateListShuffleAndSum(1, 10, 1))
	fmt.Println(CreateListShuffleAndSum(1, 10, 2))
	fmt.Println(CreateListShuffleAndSum(1, 10, 3))
	fmt.Println(CreateListShuffleAndSum(1, 10, 4))
	fmt.Println(CreateListShuffleAndSum(1, 10, 5))
	fmt.Println(CreateListShuffleAndSum(1, 10, 6))
	fmt.Println(CreateListShuffleAndSum(1, 10, 7))
	fmt.Println(CreateListShuffleAndSum(1, 10, 8))
	fmt.Println(CreateListShuffleAndSum(1, 10, 9))
	fmt.Println(CreateListShuffleAndSum(1, 10, 10))

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
