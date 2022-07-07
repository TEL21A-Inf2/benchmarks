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
// Sortiert diese Liste mit InsertionSort() und liefert die Zeit zurück.
func DurationInsertionSort(length, max int) time.Duration {
	list := randomlist.Ints(length, max)
	startTime := time.Now()
	InsertionSort(list)
	endTime := time.Now()
	return endTime.Sub(startTime)
}

// Erzeugt eine Zufallsliste mit length Elementen im Intervall [0,max].
// Sortiert diese Liste mit SelectionSort() und liefert die Zeit zurück.
func DurationSelectionSort(length, max int) time.Duration {
	list := randomlist.Ints(length, max)
	startTime := time.Now()
	SelectionSort(list)
	endTime := time.Now()
	return endTime.Sub(startTime)
}

// Erzeugt eine Zufallsliste mit length Elementen im Intervall [0,max].
// Sortiert diese Liste mit BubbleSort() und liefert die Zeit zurück.
func DurationBubbleSort(length, max int) time.Duration {
	list := randomlist.Ints(length, max)
	startTime := time.Now()
	BubbleSort(list)
	endTime := time.Now()
	return endTime.Sub(startTime)
}

// Erzeugt eine Zufallsliste mit length Elementen im Intervall [0,max].
// Sortiert diese Liste mit MergeSort() und liefert die Zeit zurück.
func DurationMergeSort(length, max int) time.Duration {
	list := randomlist.Ints(length, max)
	startTime := time.Now()
	MergeSort(list)
	endTime := time.Now()
	return endTime.Sub(startTime)
}

// Erzeugt eine Zufallsliste mit length Elementen im Intervall [0,max].
// Sortiert diese Liste mit QuickSort() und liefert die Zeit zurück.
func DurationQuickSort(length, max int) time.Duration {
	list := randomlist.Ints(length, max)
	startTime := time.Now()
	QuickSort(list)
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
	{
		var avgDuration time.Duration
		for i := 0; i < iterations; i++ {
			duration := DurationSelectionSort(max, count)
			avgDuration += duration
		}
		avgDuration /= time.Duration(iterations)
		fmt.Printf("  Durchschnittliche Dauer von SelectionSort %v\n", avgDuration)
	}
	{
		var avgDuration time.Duration
		for i := 0; i < iterations; i++ {
			duration := DurationBubbleSort(max, count)
			avgDuration += duration
		}
		avgDuration /= time.Duration(iterations)
		fmt.Printf("  Durchschnittliche Dauer von BubbleSort %v\n", avgDuration)
	}
	{
		var avgDuration time.Duration
		for i := 0; i < iterations; i++ {
			duration := DurationMergeSort(max, count)
			avgDuration += duration
		}
		avgDuration /= time.Duration(iterations)
		fmt.Printf("  Durchschnittliche Dauer von MergeSort %v\n", avgDuration)
	}
	{
		var avgDuration time.Duration
		for i := 0; i < iterations; i++ {
			duration := DurationQuickSort(max, count)
			avgDuration += duration
		}
		avgDuration /= time.Duration(iterations)
		fmt.Printf("  Durchschnittliche Dauer von QuickSort %v\n", avgDuration)
	}
}
