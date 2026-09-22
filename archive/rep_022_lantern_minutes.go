//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 022 · ch5 structs · difficulty 3/10 · ~4 min         ║
// ║ shape: FIX · tags: methods-value-receiver                     ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  A campsite estimates how long each lantern lasts.   ║
// ║ TASK                                                        ║
// ║ 1. This program compiles but gives the wrong estimates.      ║
// ║    Each unit of fuel gives 3 minutes; zero or less gives 0.  ║
// ║ 2. Name the bug, then fix the method body only.              ║
// ║    Keep the receiver, signature, struct and main unchanged.  ║
// ║ EXPECTED OUTPUT                                             ║
// ║   0                                                         ║
// ║   18                                                        ║
// ╚════════════════════════════════════════════════════════════════╝

type lantern struct {
	fuel int
}

func (l lantern) minutes() int {
	if l.fuel <= 0 {
		return 0
	}
	return l.fuel * 3
}

func main() {
	empty := lantern{}
	full := lantern{fuel: 6}
	fmt.Println(empty.minutes())
	fmt.Println(full.minutes())
}

/*
Verdict: pass, notes. Code fix clean on first check; output 0, 18 matches.
gofmt and go vet clean. Changed lantern{}.fuel to l.fuel unaided.
Explanation needed clarification: lantern{} is a new zero-valued lantern,
not a fieldless struct; receiver gets a copy of the whole lantern, not just 6.
Matt correctly traced the guard and 6 * 3, then confirmed understanding.
methods-value-receiver: revisit -> shaky; copy model still needs independent evidence.
*/
