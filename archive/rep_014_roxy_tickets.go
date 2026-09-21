//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 014 · ch3/ch4 · difficulty 3/10 · ~4 min              ║
// ║ shape: FIX · tags: guard-clauses, if-syntax                    ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The Roxy cinema has just repriced its tickets. The    ║
// ║ box office till compiles and runs, and charges the wrong       ║
// ║ people the wrong money.                                        ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. The Roxy's price list, in words:                            ║
// ║      · under 16            $7                                  ║
// ║      · 65 and over         $9                                  ║
// ║      · everyone else       $12                                 ║
// ║ 2. Fix ticketPrice so it charges that. Don't touch main.       ║
// ║ 3. There is more than one thing wrong with it.                 ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║   age 8: $7                                                    ║
// ║   age 15: $7                                                   ║
// ║   age 16: $12                                                  ║
// ║   age 40: $12                                                  ║
// ║   age 64: $12                                                  ║
// ║   age 65: $9                                                   ║
// ║   age 80: $9                                                   ║
// ╚════════════════════════════════════════════════════════════════╝

// ticketPrice returns the price in dollars for one seat at the Roxy.
func ticketPrice(age int) int {
	switch {
	case age < 16:
		return 7
	case age >= 65:
		return 9
	default:
		return 12
	}
}

func main() {
	fmt.Printf("age 8: $%d\n", ticketPrice(8))
	fmt.Printf("age 15: $%d\n", ticketPrice(15))
	fmt.Printf("age 16: $%d\n", ticketPrice(16))
	fmt.Printf("age 40: $%d\n", ticketPrice(40))
	fmt.Printf("age 64: $%d\n", ticketPrice(64))
	fmt.Printf("age 65: $%d\n", ticketPrice(65))
	fmt.Printf("age 80: $%d\n", ticketPrice(80))
}

// ── VERDICT: pass, notes ───────────────────────────────────────────
// All 7 lines match. gofmt clean, go vet clean.
// Both boundaries right: 16 → $12 and 65 → $9, which is where the
// planted `<=` was aimed.
// Note: solved by replacing the if-chain with a tagless switch rather
// than diagnosing the two planted bugs, so the boundary evidence is
// "wrote it right", not "spotted it wrong". if-syntax untested here;
// switch-tagless is what actually got exercised.
// One nudge: wrote the tagless cases under a tagged `switch age`,
// giving `cannot convert age < 16 (untyped bool value) to type int`.
// Landed it himself once pointed at the word "bool" — a case is a
// question, not a value. Not python mind; two Go switch forms crossed.
