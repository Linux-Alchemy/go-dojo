//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 011 · ch3 conditionals · difficulty 3/10 · ~4 min     ║
// ║ shape: FIX · tags: no-truthiness, logical-ops                  ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The pub quiz scoring script was ported from Python    ║
// ║ over a pint. It doesn't compile, and once it does, it scores   ║
// ║ some rounds wrong.                                             ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ HOUSE RULES                                                    ║
// ║  - A blank answer sheet (score 0) costs the team 2 points.     ║
// ║  - Playing the joker doubles the round, but only if the team   ║
// ║    scored at least 5. Otherwise the joker does nothing.        ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. Make it compile.                                            ║
// ║ 2. Make it follow the house rules.                             ║
// ║ 3. Don't touch report or main.                                 ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║  Quizteama Aguilera (joker true): 14 points                    ║
// ║  Les Quizerables (joker true): 3 points                        ║
// ║  Agatha Quiztie (joker false): 6 points                        ║
// ║  The Blank Sheets (joker true): -2 points                      ║
// ╚════════════════════════════════════════════════════════════════╝

// roundTotal returns the points a team gets for one round.
func roundTotal(score int, joker bool) int {
	if score == 0 {
		return -2
	}
	if joker && score >= 5 {
		return score * 2
	}
	return score
}

// report prints one team's result for the round.
func report(team string, score int, joker bool) {
	fmt.Printf("%s (joker %t): %d points\n", team, joker, roundTotal(score, joker))
}

func main() {
	report("Quizteama Aguilera", 7, true)
	report("Les Quizerables", 3, true)
	report("Agatha Quiztie", 6, false)
	report("The Blank Sheets", 0, true)
}

/*
VERDICT: clean

- Compile bug: `!score` → `score == 0`. Go has no truthiness, so an int has to be compared to something.
- Logic bug: `||` → `&&`. The joker needs both conditions, and Les Quizerables and Agatha Quiztie both show it.
- All four lines matched first run. gofmt and vet clean. No flags, and no redundant parens this time.
*/
