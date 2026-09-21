package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 017 · ch4 functions · difficulty 3/10  ·  ~5 min      ║
// ║ shape: COMPLETE · tags: higher-order, guard-clauses            ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  Quiz night at the Crown. The quizmaster changes how   ║
// ║ he scores the two rounds every single week, and is not to be   ║
// ║ reasoned with.                                                 ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. Fill in applyRule so it scores a team's two rounds using    ║
// ║    whichever rule it was handed. It must not care which one.   ║
// ║ 2. A team that missed a round has a negative score for it.     ║
// ║    They score 0 for the night, whatever the rule would say.    ║
// ║ 3. Don't touch main, bestOf or combined.                       ║
// ║                                                                ║
// ║ It compiles as it stands, and it is wrong as it stands.        ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║   19                                                           ║
// ║   31                                                           ║
// ║   0                                                            ║
// ║   0                                                            ║
// ║   7                                                            ║
// ╚════════════════════════════════════════════════════════════════╝

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
	return 0
}

func main() {
	fmt.Println(applyRule(12, 19, bestOf))
	fmt.Println(applyRule(12, 19, combined))
	fmt.Println(applyRule(-1, 19, bestOf))
	fmt.Println(applyRule(12, -1, combined))
	fmt.Println(applyRule(0, 7, bestOf))
}
