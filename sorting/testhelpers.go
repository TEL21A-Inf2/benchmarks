package sorting

import (
	"sort"
	"testing"
)

// Hilfsfunktion: Wendet eine Funktion auf die Liste an, prüft,
// ob die Liste anschließend sortiert ist und löst ggf. den Fehler aus.
func testSortAlgorithm(sortFunc func([]int), list []int, name string, t *testing.T) {
	sortFunc(list)
	if !sort.IntsAreSorted(list) {
		t.Errorf("%s: %v sollte sortiert sein, ist es aber nicht.", name, list)
	}
}
