package sorting

// Implementierung von HeapSort.
func HeapSort(list []int) {
	if len(list) == 0 {
		return
	}
	for i := range list {
		currentPos := i
		parentPos := (currentPos - 1) / 2

		for currentPos > 0 && list[currentPos] > list[parentPos] {
			list[parentPos], list[currentPos] = list[currentPos], list[parentPos]
			currentPos = parentPos
			parentPos = (currentPos - 1) / 2
		}
	}

	size := len(list) - 1
	for size > 0 {
		list[0], list[size] = list[size], list[0]
		size--

		currentPos := 0
		doneSomething := true
		for currentPos < size && doneSomething {
			leftChildPos, rightChildPos := 2*currentPos+1, 2*currentPos+2
			doneSomething = false

			if rightChildPos <= size {
				if list[leftChildPos] < list[rightChildPos] {
					if list[currentPos] < list[rightChildPos] {
						list[currentPos], list[rightChildPos] = list[rightChildPos], list[currentPos]
						currentPos = rightChildPos
						doneSomething = true
					}
				} else {
					if list[currentPos] < list[leftChildPos] {
						list[currentPos], list[leftChildPos] = list[leftChildPos], list[currentPos]
						currentPos = leftChildPos
						doneSomething = true
					}
				}
			} else if leftChildPos <= size {
				if list[currentPos] < list[leftChildPos] {
					list[currentPos], list[leftChildPos] = list[leftChildPos], list[currentPos]
					currentPos = leftChildPos
					doneSomething = true
				}
			}
		}
	}
}
