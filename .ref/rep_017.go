package main

import "fmt"

// bestOf returns the higher of two round scores.
func bestOf(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// combined returns the two round scores added together.
func combined(x, y int) int {
	return x + y
}

// applyRule scores one team's two rounds using whichever rule the
// quizmaster picked tonight.
func applyRule(a, b int, rule func(int, int) int) int {
	if a < 0 || b < 0 {
		return 0
	}
	return rule(a, b)
}

func main() {
	fmt.Println(applyRule(12, 19, bestOf))
	fmt.Println(applyRule(12, 19, combined))
	fmt.Println(applyRule(-1, 19, bestOf))
	fmt.Println(applyRule(12, -1, combined))
	fmt.Println(applyRule(0, 7, bestOf))
}
