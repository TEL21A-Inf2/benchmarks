package branchprediction

import (
	"fmt"
	"sort"
	"time"

	"github.com/tel21a-inf2/benchmarks/randomlist"
	"github.com/tel21a-inf2/benchmarks/timing"
)

// Erwartet eine Liste von Zahlen und eine Zahl max.
// Liefert die Summe aller Elemente der Liste, die größer als max sind.
func SumGreater(list []int, max int) int {
	result := 0
	for _, el := range list {
		if el > max {
			result += el
		}
	}
	return result
}

// Erzeugt eine Zufallsliste der Länge length mit Werten zwischen 0 und max.
// Summiert dann die Elemente auf, die größer als threshold sind.
// Liefert die Zeit, die das Aufsummieren gedauert hat.
func CreateListAndSum(max, length int) time.Duration {
	list := randomlist.Ints(length, max)
	duration := timing.Duration(func() { SumGreater(list, max/2) })
	return duration
}

// Erzeugt eine Zufallsliste der Länge length mit Werten zwischen 0 und max.
// Summiert dann die Elemente auf, die größer als threshold sind.
// Liefert die Zeit, die das Aufsummieren gedauert hat.
// Sortiert die Liste vor dem Aufsummieren.
func CreateListSortAndSum(max, length int) time.Duration {
	list := randomlist.Ints(length, max)
	sort.Ints(list)
	duration := timing.Duration(func() { SumGreater(list, max/2) })
	return duration
}

// Führt den eigentlichen Benchmark aus.
// Erwartet Anzahl und Höchstgrenze für die Elemente der Liste, sowie die Anzahl der Wiederholungen.
// Summiert alles, was größer ist als count / 2.
// Ausgeführt wird jeweils einmal CreateListAndSum() und CreateListSortAndSum().
// Gibt die Summe sowie die durchschnittliche Dauer auf die Konsole aus.
func RunBenchmark(max, count, iterations int) {

	var avgDurationSorted time.Duration
	for i := 0; i < iterations; i++ {
		durationSorted := CreateListSortAndSum(max, count)
		avgDurationSorted += durationSorted
	}
	avgDurationSorted /= time.Duration(iterations)

	var avgDurationShuffled time.Duration
	for i := 0; i < iterations; i++ {
		durationShuffled := CreateListAndSum(max, count)
		avgDurationShuffled += durationShuffled
	}
	avgDurationShuffled /= time.Duration(iterations)

	fmt.Printf("%v Elemente zwischen 0 und %v, summiert werden Elemente über %v:\n", count, max, max/2)
	fmt.Printf("  Durchschnittliche Dauer bei sortierter Liste: %v\n", avgDurationSorted)
	fmt.Printf("  Durchschnittliche Dauer bei unsortierter Liste: %v\n\n", avgDurationShuffled)

}
