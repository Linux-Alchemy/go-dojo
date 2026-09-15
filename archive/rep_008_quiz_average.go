//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 008 · ch1/ch4 · difficulty 3/10 · ~3 min              ║
// ║ shape: FIX · tags: type-conversion, int-division, grouped-params║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The pub quiz league table shows each team's average   ║
// ║ points per round. Teams are complaining the table is stingy.   ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. This compiles, but it prints the wrong averages. Find the   ║
// ║    bug and fix it.                                             ║
// ║ 2. A team that played no rounds must still average 0.00.      ║
// ║ 3. Don't touch main or the Printf lines.                       ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║  Quizzly Bears:   7.83                                         ║
// ║  Les Quizerables: 10.00                                        ║
// ║  Agatha Quiztie:  7.25                                         ║
// ║  No-Shows:        0.00                                         ║
// ╚════════════════════════════════════════════════════════════════╝

// averageScore returns a team's mean points per round.
func averageScore(total, rounds float64) float64 {
	if rounds == 0 {
		return 0
	}
	return float64(total / rounds)
}

func main() {
	fmt.Printf("Quizzly Bears:   %.2f\n", averageScore(47, 6))
	fmt.Printf("Les Quizerables: %.2f\n", averageScore(60, 6))
	fmt.Printf("Agatha Quiztie:  %.2f\n", averageScore(29, 4))
	fmt.Printf("No-Shows:        %.2f\n", averageScore(0, 0))
}

// ── VERDICT: pass, notes ──
// Output matched first run; gofmt + vet clean. Explanation of the bug was spot on.
// Flag 1: float64(total / rounds) is now a no-op conversion; both are already float64.
// Flag 2: fixed by changing the signature to float64 params. Works here only because main passes
//         untyped constants; a caller holding int variables would no longer compile.
//         Counts stay int; convert the inputs inside: float64(total) / float64(rounds).
