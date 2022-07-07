package sorting

// Implementierung von InsertionSort.
func SelectionSort(list []int) {
	for i := range list {
		minPos := i
		for j, v := range list[i+1:] {
			if v < list[minPos] {
				minPos = j + i + 1
			}
		}
		list[i], list[minPos] = list[minPos], list[i]
	}
}
