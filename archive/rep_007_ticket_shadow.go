//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 007 · ch1 variables · difficulty 3/10 · ~3 min         ║
// ║ shape: PREDICT · tags: short-decl, shadowing                    ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  A tiny theatre tracks its remaining tickets.          ║
// ║ Three printouts report the count as this program runs.         ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                          ║
// ║ 1. Write your predicted three output lines in a comment        ║
// ║    above main BEFORE running the program.                     ║
// ║ 2. Leave the executable code alone. Say check when ready.      ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT  Hidden for this PREDICT; reference verified.  ║
// ╚════════════════════════════════════════════════════════════════╝

// Your prediction:
//inside: 5
//after: 8
//final: 6

func main() {
	tickets := 8
	if tickets > 0 {
		tickets := tickets - 3
		fmt.Println("inside:", tickets)
	}
	fmt.Println("after:", tickets)
	tickets = tickets - 2
	fmt.Println("final:", tickets)
}

// ── VERDICT: clean ──
// Predicted inside: 5 / after: 8 / final: 6, matched on first run.
// Spotted that the inner := makes a new tickets that dies at the closing brace.
// gofmt + vet clean. No flags.
