package main

import "fmt"

func roundTotal(score int, joker bool) int {
	if score == 0 {
		return -2
	}
	if joker && score >= 5 {
		return score * 2
	}
	return score
}

func report(team string, score int, joker bool) {
	fmt.Printf("%s (joker %t): %d points\n", team, joker, roundTotal(score, joker))
}

func main() {
	report("Quizteama Aguilera", 7, true)
	report("Les Quizerables", 3, true)
	report("Agatha Quiztie", 6, false)
	report("The Blank Sheets", 0, true)
}
