package branchprediction

import (
	"fmt"
	"sort"
	"time"

	"github.com/tel21a-inf2/benchmarks/randomlist"
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

// Erzeugt eine Zufallsliste der Länge length mit Werten zwischen 0 und max.
// Summiert dann die Elemente auf, die größer als threshold sind.
// Liefert diese Summe und die Zeit, die das Aufsummieren gedauert hat.
func CreateListAndSum(max, length int) (int, time.Duration) {
	list := randomlist.Ints(length, max)
	var sum int
	duration := timing.Duration(func() { sum = SumGreater(list, max/2) })
	return sum, duration
}

// Erzeugt eine Zufallsliste der Länge length mit Werten zwischen 0 und max.
// Summiert dann die Elemente auf, die größer als threshold sind.
// Liefert diese Summe und die Zeit, die das Aufsummieren gedauert hat.
// Sortiert die Liste vor dem Aufsummieren.
func CreateListSortAndSum(max, length int) (int, time.Duration) {
	list := randomlist.Ints(length, max)
	sort.Ints(list)
	var sum int
	duration := timing.Duration(func() { sum = SumGreater(list, max/2) })
	return sum, duration
}

// Führt den eigentlichen Benchmark aus.
// Erwartet Anzahl und Höchstgrenze für die Elemente der Liste.
// Summiert alles, was größer ist als count / 2.
// Ausgeführt wird jeweils einmal CreateListAndSum() und CreateListSortAndSum().
// Gibt die Summe sowie die Dauer auf die Konsole aus.
func RunBenchmark(max, count int) {

	_, durationSorted := CreateListSortAndSum(max, count)
	_, durationShuffled := CreateListAndSum(max, count)

	fmt.Printf("%v Elemente zwischen 0 und %v, summiert werden Elemente über %v:\n", count, max, count/2)

	fmt.Printf("  Dauer bei sortierter Liste: %v\n", durationSorted)
	fmt.Printf("  Dauer bei unsortierter Liste: %v\n\n", durationShuffled)

}
