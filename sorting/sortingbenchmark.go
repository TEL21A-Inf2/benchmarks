package sorting

import (
	"fmt"
	"sort"
	"time"

	"github.com/tel21a-inf2/benchmarks/randomlist"
)

// Erzeugt eine Zufallsliste mit length Elementen im Intervall [0,max].
// Sortiert diese Liste mit sort.Ints() und liefert die Zeit zurück.
func DurationSortInts(length, max int) time.Duration {
	list := randomlist.Ints(length, max)
	startTime := time.Now()
	sort.Ints(list)
	endTime := time.Now()
	return endTime.Sub(startTime)
}

// Erzeugt eine Zufallsliste mit length Elementen im Intervall [0,max].
// Sortiert diese Liste mit IsertionSort() und liefert die Zeit zurück.
func DurationInsertionSort(length, max int) time.Duration {
	list := randomlist.Ints(length, max)
	startTime := time.Now()
	InsertionSort(list)
	endTime := time.Now()
	return endTime.Sub(startTime)
}

// Führt den Sortier-Benchmark aus.
// Erwartet Anzahl und Höchstgrenze für die Elemente der Liste, sowie die Anzahl der Wiederholungen.
// Es wird jede der obigen Sortierfunktionen entsprechend der Wiederholungen ausgeführt
// und die Durchschnittsdauern ausgegeben.
func RunBenchmark(max, count, iterations int) {
	fmt.Println("Sortier-Benchmark:")
	fmt.Printf("  %v Elemente zwischen 0 und %v:\n", count, max)
	{
		var avgDuration time.Duration
		for i := 0; i < iterations; i++ {
			duration := DurationSortInts(max, count)
			avgDuration += duration
		}
		avgDuration /= time.Duration(iterations)
		fmt.Printf("  Durchschnittliche Dauer von sort.Ints %v\n", avgDuration)
	}
	{
		var avgDuration time.Duration
		for i := 0; i < iterations; i++ {
			duration := DurationInsertionSort(max, count)
			avgDuration += duration
		}
		avgDuration /= time.Duration(iterations)
		fmt.Printf("  Durchschnittliche Dauer von InsertionSort %v\n", avgDuration)
	}
}
