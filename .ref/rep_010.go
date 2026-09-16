package main

import "fmt"

// avgSpeed returns average speed in km/h.
func avgSpeed(distanceKm int, minutes int) float64 {
	hours := float64(minutes) / 60
	return float64(distanceKm) / hours
}

func main() {
	km, mins := 100, 240
	fmt.Printf("%d km in %d min: %.1f km/h\n", km, mins, avgSpeed(km, mins))

	km, mins = 30, 90
	fmt.Printf("%d km in %d min: %.1f km/h\n", km, mins, avgSpeed(km, mins))

	km, mins = 12, 45
	fmt.Printf("%d km in %d min: %.1f km/h\n", km, mins, avgSpeed(km, mins))
}
