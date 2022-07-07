package timing

import "time"

// Erwartet eine Funktion, die keine Argumente erwartet und auch nichts liefert.
// Liefert die Zeit, die die Ausführung dieser Funktion benötigt hat.
func Duration(f func()) time.Duration {
	startTime := time.Now()
	f()
	endTime := time.Now()
	return endTime.Sub(startTime)
}
