package lists

// Erwartet eine Liste aus Zahlen list und eine Zahl lenght.
// Liefert die Teilliste von list mit der Länge n, deren Elemente die höchste Summe haben.
func GreatestSublist(list []int, length int) []int {
	if length > len(list) {
		return make([]int, 0)
	}

	resultPos := 0

	for i := 0; i <= len(list)-length; i++ {

		currentSegment := list[i : i+length]
		currentSegmentSum := 0
		for _, el := range currentSegment {
			currentSegmentSum += el
		}

		resultSegment := list[resultPos : resultPos+length]
		resultSegmentSum := 0
		for _, el := range resultSegment {
			resultSegmentSum += el
		}

		if currentSegmentSum > resultSegmentSum {
			resultPos = i
		}
	}
	return list[resultPos : resultPos+length]
}
