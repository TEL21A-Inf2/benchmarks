package sorting

// Implementierung von MergeSort.
func MergeSort(list []int) {
	length := len(list)
	if length <= 1 {
		return
	}
	middle := length / 2
	left, right := list[:middle], list[middle:]
	MergeSort(left)
	MergeSort(right)

	result := make([]int, length)
	i := 0
	for ; len(left) > 0 && len(right) > 0; i++ {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
	}
	for ; len(left) > 0; i++ {
		result[i] = left[0]
		left = left[1:]
	}
	for ; len(right) > 0; i++ {
		result[i] = right[0]
		right = right[1:]
	}
	copy(list, result)
}
