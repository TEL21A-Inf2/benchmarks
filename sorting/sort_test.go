package sorting

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	l1 := []int{5, 235, 2, 3, 42, 1, 107, 7, 12}
	l2 := []int{1, 2, 3, 4, 5}
	l3 := []int{}

	for _, l := range [][]int{l1, l2, l3} {
		testSortAlgorithm(InsertionSort, l, "InsertionSort", t)
	}
}
