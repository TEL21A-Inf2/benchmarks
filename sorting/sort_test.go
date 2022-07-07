package sorting

import (
	"sort"
	"testing"
)

func TestSortInts(t *testing.T) {
	l1 := []int{5, 235, 2, 3, 42, 1, 107, 7, 12}
	l2 := []int{1, 2, 3, 4, 5}
	l3 := []int{}

	for _, l := range [][]int{l1, l2, l3} {
		testSortAlgorithm(sort.Ints, l, "sort.Ints", t)
	}
}

func TestInsertionSort(t *testing.T) {
	l1 := []int{5, 235, 2, 3, 42, 1, 107, 7, 12}
	l2 := []int{1, 2, 3, 4, 5}
	l3 := []int{}

	for _, l := range [][]int{l1, l2, l3} {
		testSortAlgorithm(InsertionSort, l, "InsertionSort", t)
	}
}

func TestSelectionSort(t *testing.T) {
	l1 := []int{5, 235, 2, 3, 42, 1, 107, 7, 12}
	l2 := []int{1, 2, 3, 4, 5}
	l3 := []int{}

	for _, l := range [][]int{l1, l2, l3} {
		testSortAlgorithm(SelectionSort, l, "SelectionSort", t)
	}
}

func TestBubbleSort(t *testing.T) {
	l1 := []int{5, 235, 2, 3, 42, 1, 107, 7, 12}
	l2 := []int{1, 2, 3, 4, 5}
	l3 := []int{}

	for _, l := range [][]int{l1, l2, l3} {
		testSortAlgorithm(BubbleSort, l, "BubbleSort", t)
	}
}

func TestMergeSort(t *testing.T) {
	l1 := []int{5, 235, 2, 3, 42, 1, 107, 7, 12}
	l2 := []int{1, 2, 3, 4, 5}
	l3 := []int{}

	for _, l := range [][]int{l1, l2, l3} {
		testSortAlgorithm(MergeSort, l, "MergeSort", t)
	}
}

func TestQuickSort(t *testing.T) {
	l1 := []int{5, 235, 2, 3, 42, 1, 107, 7, 12}
	l2 := []int{1, 2, 3, 4, 5}
	l3 := []int{}

	for _, l := range [][]int{l1, l2, l3} {
		testSortAlgorithm(QuickSort, l, "QuickSort", t)
	}
}
