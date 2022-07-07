package lists

import "fmt"

func ExampleGreatestSublist() {
	l1 := []int{1, 5, 2, 3, 7, 5, 3}

	fmt.Println(GreatestSublist(l1, 2))
	fmt.Println(GreatestSublist(l1, 3))
	fmt.Println(GreatestSublist(l1, 4))
	fmt.Println(GreatestSublist(l1, 5))
	fmt.Println(GreatestSublist(l1, 6))
	fmt.Println(GreatestSublist(l1, 7))
	fmt.Println(GreatestSublist(l1, 8))
	fmt.Println()

	l2 := []int{0, 0, 0, 10, 0, 6, 9, 2, 4, 9}
	fmt.Println(GreatestSublist(l2, 1))
	fmt.Println(GreatestSublist(l2, 2))
	fmt.Println(GreatestSublist(l2, 3))
	fmt.Println(GreatestSublist(l2, 4))
	fmt.Println(GreatestSublist(l2, 5))

	// Output:
	// [7 5]
	// [3 7 5]
	// [3 7 5 3]
	// [5 2 3 7 5]
	// [5 2 3 7 5 3]
	// [1 5 2 3 7 5 3]
	// []
	//
	// [10]
	// [6 9]
	// [6 9 2]
	// [10 0 6 9]
	// [6 9 2 4 9]
}
