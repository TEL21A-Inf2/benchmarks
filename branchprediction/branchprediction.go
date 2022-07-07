package branchprediction

import "math/rand"

// Erwartet eine Liste von Zahlen und eine Zahl t.
// Liefert die Summe aller Elemente der Liste, die größer als t sind.
func SumGreater(list []int, x int) int {
	result := 0
	for _, el := range list {
		if el > x {
			result += el
		}
	}
	return result
}

// Erzeugt eine Liste der Länge n, die die ersten n Vielfachen von x enthält.
func Multiples(x, n int) []int {
	result := make([]int, n)
	currentValue := x
	for i := 0; i < n; i++ {
		result[i] = currentValue
		currentValue += x
	}
	return result
}

// Erzeugt eine Liste mit Multiples(x,n) und summiert
// dann die Elemente auf, die größer als t sind.
func CreateListAndSum(x, n, t int) int {
	list := Multiples(x, n)
	sum := SumGreater(list, t)
	return sum
}

// Erzeugt eine Liste mit Multiples(x,n) und summiert
// dann die Elemente auf, die größer als t sind.
// Bringt die Liste vor dem Aufsummieren in eine zufällige Reihenfolge.
func CreateListShuffleAndSum(x, n, t int) int {
	list := Multiples(x, n)
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	sum := SumGreater(list, t)
	return sum
}
