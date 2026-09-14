//go:build ignore

package main

import "fmt"

// ╔══════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 002  ·  ch1–2 variables, fmt  ·  difficulty 2/10  ·  ~2m ║
// ║ shape: PREDICT  ·  tags: int-division, type-conversion,           ║
// ║                         printf-verbs                              ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  Pub quiz night. Three pizzas, eight slices each, seven   ║
// ║ teams. The landlord wants to know how it shares out.              ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ TASK                                                              ║
// ║  1. DON'T RUN IT YET.                                             ║
// ║  2. Write the four lines you think it prints in the PREDICTION    ║
// ║     comment below.                                                ║
// ║  3. Then run it and compare. Say `check` either way.              ║
// ║  HINT  two of the four lines are where Python would disagree.     ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                   ║
// ║  (that's your job this time)                                      ║
// ╚══════════════════════════════════════════════════════════════════╝

// PREDICTION:
//3
//3.4
//int float64
//24 slices for 7 teams

func main() {
	pizzas := 3
	teams := 7
	slices := pizzas * 8
	each := float64(slices) / float64(teams)

	fmt.Println(slices / teams)
	fmt.Printf("%.1f\n", each)
	fmt.Printf("%T %T\n", slices, each)
	fmt.Println(slices, "slices for", teams, "teams")
}

// ── VERDICT: pass, notes ─────────────────────────────────────────
// Prediction matches output. Ran before predicting (asked why line 36
// printed 3) and inlay hints were showing the types, so not a blind
// predict. int-division explained: int / int truncates, like Python //.
// - Style nit: `//3` → `// 3`, space after // is Go convention.
