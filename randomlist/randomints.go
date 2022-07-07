package randomlist

import "math/rand"

// Liefert eine Liste der Länge length mit Zufallszahlen, die zwischen 0 und max liegen.
func Ints(length, max int) []int {
	result := make([]int, length)
	for i := 0; i < length; i++ {
		result[i] = rand.Intn(max)
	}
	return result
}
