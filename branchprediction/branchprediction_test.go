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
	for i := 0; i <= 10; i++ {
		result, _ := CreateListAndSum(1, 10, i)
		fmt.Println(result)
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

func ExampleCreateListShuffleAndSum() {
	for i := 0; i <= 10; i++ {
		result, _ := CreateListShuffleAndSum(1, 10, i)
		fmt.Println(result)
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
