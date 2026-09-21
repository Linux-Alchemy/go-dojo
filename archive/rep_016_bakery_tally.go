//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 016 · ch4 functions · difficulty 3/10  ·  ~4 min      ║
// ║ shape: FIX · tags: closures, shadowing                         ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The bakery counts what goes out of the door. It       ║
// ║ compiles, it runs, go vet is perfectly happy with it, and the  ║
// ║ day's totals are nonsense.                                     ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. tally() should hand back a function that keeps a RUNNING    ║
// ║    total: each call adds to what came before. Separate         ║
// ║    tallies must not interfere with each other.                 ║
// ║ 2. Find why it doesn't, and fix it. One line is wrong.         ║
// ║ 3. Don't touch main.                                           ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║   12                                                           ║
// ║   20                                                           ║
// ║   50                                                           ║
// ║   5                                                            ║
// ║   51                                                           ║
// ╚════════════════════════════════════════════════════════════════╝

// tally hands back a running-total function for one counter.
func tally() func(int) int {
	sold := 0
	return func(n int) int {
		sold += n
		return sold
	}
}

func main() {
	rolls := tally()
	fmt.Println(rolls(12))
	fmt.Println(rolls(8))
	fmt.Println(rolls(30))

	pies := tally()
	fmt.Println(pies(5))
	fmt.Println(rolls(1))
}

// ── VERDICT: clean ─────────────────────────────────────────────────
// `sold := sold + n` → `sold += n`, first run, no nudges. All 5 lines
// match. gofmt clean, go vet clean. No flags.
// Spotted that a new variable declared inside the inner function is
// reborn on every call, so the enclosed one never moves — despite the
// program compiling and vetting clean, which was the point of the rep.
// closures: shaky → SOLID (clean PREDICT 015 + clean FIX 016).
// shadowing: shaky → SOLID (clean PREDICT 007 + clean FIX 016).
