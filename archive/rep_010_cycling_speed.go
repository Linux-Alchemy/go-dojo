//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 010 · ch1/ch4 · difficulty 3/10 · ~4 min              ║
// ║ shape: TRANSLATE · tags: type-conversion, int-division         ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  A cycling app is being ported from Python to Go.      ║
// ║ The ride log stores whole kilometres and whole minutes.        ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. Translate the Python below into the body of avgSpeed.       ║
// ║ 2. The signature is fixed: two ints in, a float64 out.         ║
// ║ 3. Don't touch main.                                           ║
// ║                                                                ║
// ║   def avg_speed(distance_km, minutes):                         ║
// ║       hours = minutes / 60                                     ║
// ║       return distance_km / hours                               ║
// ║                                                                ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║  100 km in 240 min: 25.0 km/h                                  ║
// ║  30 km in 90 min: 20.0 km/h                                    ║
// ║  12 km in 45 min: 16.0 km/h                                    ║
// ╚════════════════════════════════════════════════════════════════╝

// avgSpeed returns average speed in km/h.
func avgSpeed(distanceKm int, minutes int) float64 {
	timeHours := (float64(minutes) / 60)
	return (float64(distanceKm) / timeHours)
}

func main() {
	km, mins := 100, 240
	fmt.Printf("%d km in %d min: %.1f km/h\n", km, mins, avgSpeed(km, mins))

	km, mins = 30, 90
	fmt.Printf("%d km in %d min: %.1f km/h\n", km, mins, avgSpeed(km, mins))

	km, mins = 12, 45
	fmt.Printf("%d km in %d min: %.1f km/h\n", km, mins, avgSpeed(km, mins))
}

/*
VERDICT: pass, notes

- Output matched all three lines first run, no nudges. gofmt and vet clean.
- Converted the int inputs before dividing, and left the signature alone. That's the 008 lesson applied without prompting.
- Dividing a float64 by the bare 60 is fine: an untyped constant takes whatever type it needs.

FLAGS
1. Redundant parens around both expressions: `(float64(minutes) / 60)` and the return. Go doesn't need them (second rep running, see 009).
*/
