package main

import "fmt"

func averageScore(total, rounds int) float64 {
	if rounds == 0 {
		return 0
	}
	return float64(total) / float64(rounds)
}

func main() {
	fmt.Printf("Quizzly Bears:   %.2f\n", averageScore(47, 6))
	fmt.Printf("Les Quizerables: %.2f\n", averageScore(60, 6))
	fmt.Printf("Agatha Quiztie:  %.2f\n", averageScore(29, 4))
	fmt.Printf("No-Shows:        %.2f\n", averageScore(0, 0))
}
