package branchprediction

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/tel21a-inf2/benchmarks/timing"
)

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
// Liefert diese Summe und die Zeit, die das Aufsummieren gedauert hat.
func CreateListAndSum(x, n, t int) (int, time.Duration) {
	list := Multiples(x, n)
	var sum int
	duration := timing.Duration(func() { sum = SumGreater(list, t) })
	return sum, duration
}

// Erzeugt eine Liste mit Multiples(x,n) und summiert
// dann die Elemente auf, die größer als t sind.
// Liefert diese Summe und die Zeit, die das Aufsummieren gedauert hat.
func CreateListShuffleAndSum(x, n, t int) (int, time.Duration) {
	list := Multiples(x, n)
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
	var sum int
	duration := timing.Duration(func() { sum = SumGreater(list, t) })
	return sum, duration
}

// Führt den eigentlichen Benchmark aus:
//
// Erzeugt eine Liste von count Vielfachen von base,
// summiert dann deren Elemente über threshold.
// Ausgeführt wird jeweils einmal CreateListAndSum() und CreateListShuffleAndSum().
// Gibt die Summe sowie die Dauer auf die Konsole aus.
func RunBenchmark(base, count, threshold int) {
	sumSorted, durationSorted := CreateListAndSum(base, count, threshold)
	sumShuffled, durationShuffled := CreateListShuffleAndSum(base, count, threshold)

	fmt.Printf("%v Vielfache von %v, summiere die Elemente > %v:\n", count, base, threshold)

	fmt.Println("  Sortierte Liste:")
	fmt.Printf("    Ergebnis: %v, Dauer: %v\n\n", sumSorted, durationSorted)

	fmt.Println("  Unsortierte Liste:")
	fmt.Printf("    Ergebnis: %v, Dauer: %v\n\n", sumShuffled, durationShuffled)

}
