package sorting

import (
	"sort"
	"testing"
)

func TestSortInts(t *testing.T) {
	for _, l := range testData() {
		testSortAlgorithm(sort.Ints, l, "sort.Ints", t)
	}
}

func TestInsertionSort(t *testing.T) {
	for _, l := range testData() {
		testSortAlgorithm(InsertionSort, l, "InsertionSort", t)
	}
}

func TestSelectionSort(t *testing.T) {
	for _, l := range testData() {
		testSortAlgorithm(SelectionSort, l, "SelectionSort", t)
	}
}

func TestBubbleSort(t *testing.T) {
	for _, l := range testData() {
		testSortAlgorithm(BubbleSort, l, "BubbleSort", t)
	}
}

func TestMergeSort(t *testing.T) {
	for _, l := range testData() {
		testSortAlgorithm(MergeSort, l, "MergeSort", t)
	}
}

func TestQuickSort(t *testing.T) {
	for _, l := range testData() {
		testSortAlgorithm(QuickSort, l, "QuickSort", t)
	}
}

func TestHeapSort(t *testing.T) {
	for _, l := range testData() {
		testSortAlgorithm(HeapSort, l, "HeapSort", t)
	}
}

// Hilfsfunktion: Liefert die Test-Listen für die Tests.
func testData() [][]int {
	l1 := []int{5, 235, 2, 3, 42, 1, 107, 7, 12}
	l2 := []int{1, 2, 3, 4, 5}
	l3 := []int{}

	return [][]int{l1, l2, l3}
}

// Hilfsfunktion: Wendet eine Funktion auf die Liste an, prüft,
// ob die Liste anschließend sortiert ist und löst ggf. den Fehler aus.
func testSortAlgorithm(sortFunc func([]int), list []int, name string, t *testing.T) {
	sortFunc(list)
	if !sort.IntsAreSorted(list) {
		t.Errorf("%s: %v sollte sortiert sein, ist es aber nicht.", name, list)
	}
}
