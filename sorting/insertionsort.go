package sorting

// Implementierung von InsertionSort.
func InsertionSort(list []int) {
	for i := range list {
		for j := i; j > 0 && list[j] < list[j-1]; j-- {
			list[j], list[j-1] = list[j-1], list[j]
		}
	}
}
