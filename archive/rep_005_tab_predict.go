//go:build ignore

package main

import "fmt"

// ╔══════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 005  ·  ch1 variables  ·  difficulty 3/10  ·  ~3 min     ║
// ║ shape: PREDICT  ·  tags: int-division, type-conversion            ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  Four mates, one $45 tab. Three ways of splitting it,     ║
// ║ and only one of them is honest.                                   ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ TASK                                                              ║
// ║  0. Inlay hints OFF first (<leader>uh). No peeking at types.      ║
// ║  1. Read main. Don't run it yet.                                  ║
// ║  2. Write your predicted output (6 lines) as a comment in the     ║
// ║     PREDICTION block below, one line per print.                   ║
// ║  3. Then say "check". I'll run it and compare.                    ║
// ╚══════════════════════════════════════════════════════════════════╝

// PREDICTION:
//11
//44
//11.25
//11.00
//11
//11.25

func main() {
	tab := 45
	mates := 4

	share := tab / mates
	exact := float64(tab) / float64(mates)
	oops := float64(tab / mates)

	fmt.Println(share)
	fmt.Println(share * mates)
	fmt.Printf("%.2f\n", exact)
	fmt.Printf("%.2f\n", oops)
	fmt.Println(oops)
	fmt.Println(exact)
}

// ── VERDICT: pass, notes ─────────────────────────────────────────
// All six lines predicted correctly, but after asking about lines 4–5.
// - float64(a / b) converts after the int division has already truncated.
// - Println on a whole-valued float64 prints 11, not 11.0 (still float64).
