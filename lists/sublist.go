package lists

import (
	"fmt"
	"time"

	"github.com/tel21a-inf2/benchmarks/randomlist"
)

// Erwartet eine Liste aus Zahlen list und eine Zahl length.
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

func GreatestSublistOptimized(list []int, length int) []int {
	if length > len(list) {
		return make([]int, 0)
	}

	resultPos := 0
	greatestSum := 0
	for _, v := range list[0:length] {
		greatestSum += v
	}
	currentSum := greatestSum

	for i := length; i < len(list); i++ {
		currentSum = currentSum - list[i-length] + list[i]
		if currentSum > greatestSum {
			greatestSum = currentSum
			resultPos = i - length + 1
		}
	}
	return list[resultPos : resultPos+length]
}

// Erzeugt eine Zufallsliste mit length Elementen im Intervall [0,max].
// Führt GreatestSublist darauf aus und liefert die Zeit zurück.
func DurationGreatestSublist(length, max, sublength int) time.Duration {
	list := randomlist.Ints(length, max)
	startTime := time.Now()
	GreatestSublist(list, sublength)
	endTime := time.Now()
	return endTime.Sub(startTime)
}

// Erzeugt eine Zufallsliste mit length Elementen im Intervall [0,max].
// Führt GreatestSublistOptimized darauf aus und liefert die Zeit zurück.
func DurationGreatestSublistOptimized(length, max, sublength int) time.Duration {
	list := randomlist.Ints(length, max)
	startTime := time.Now()
	GreatestSublistOptimized(list, sublength)
	endTime := time.Now()
	return endTime.Sub(startTime)
}

// Führt den Sublist-Benchmark aus.
// Erwartet Anzahl und Höchstgrenze für die Elemente der Liste, die Länge der Subliste
// sowie die Anzahl der Wiederholungen.
// Führt entsprechend oft GreatestSublist bzw. die optimierte Fassung aus und gibt
// die durchschnittliche Laufzeit auf der Konsole aus.
func RunBenchmark(max, count, sublength, iterations int) {
	fmt.Println("Benchmark für GreatestSublist")
	fmt.Printf("  %v Elemente zwischen 0 und %v\n", count, max)
	{
		var avgDuration time.Duration
		for i := 0; i < iterations; i++ {
			duration := DurationGreatestSublist(max, count, sublength)
			avgDuration += duration
		}
		avgDuration /= time.Duration(iterations)
		fmt.Printf("  Durchschnittliche Dauer bei naiver Variante: %v\n", avgDuration)
	}
	{
		var avgDuration time.Duration
		for i := 0; i < iterations; i++ {
			duration := DurationGreatestSublistOptimized(max, count, sublength)
			avgDuration += duration
		}
		avgDuration /= time.Duration(iterations)
		fmt.Printf("  Durchschnittliche Dauer bei optimierter Variante: %v\n", avgDuration)
	}
	fmt.Println()
}
