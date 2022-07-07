package sorting

// Implementierung von QuickSort.
func QuickSort(list []int) {
	length := len(list)
	if length <= 1 {
		return
	}
	pivot := list[0]

	leftPos, rightPos := 1, length-1
	for leftPos <= rightPos {
		for leftPos <= rightPos && list[leftPos] <= pivot {
			leftPos++
		}
		for leftPos <= rightPos && list[rightPos] > pivot {
			rightPos--
		}
		if leftPos < rightPos && list[leftPos] > list[rightPos] {
			list[leftPos], list[rightPos] = list[rightPos], list[leftPos]
		}
	}
	list[rightPos], list[0] = list[0], list[rightPos]

	left, right := list[:rightPos], list[leftPos:]
	QuickSort(left)
	QuickSort(right)
}
