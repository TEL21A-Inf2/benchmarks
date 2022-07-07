package sorting

// Implementierung von InsertionSort.
func BubbleSort(list []int) {
	for i := range list {
		for j, v := range list[:len(list)-i-1] {
			if v > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}
