package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 024 · ch5 structs · difficulty 3/10 · ~4 min          ║
// ║ shape: FIX · tags: embedded-structs, short-decl                ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  Pub quiz scoreboard. Any team scoring 7 or more in a  ║
// ║          round gets a 3-point bonus.                           ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║  This doesn't compile. Two bugs, one error each.               ║
// ║  1. Fix both. Leave the embedding in place: entry must still   ║
// ║     embed team, and main must still print e.name directly.    ║
// ║  2. Before you type the fix, write one comment line above each ║
// ║     bug naming what was wrong.                                 ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║   Quizzly Bears captained by Ada                               ║
// ║   round 3: 10 points                                           ║
// ╚════════════════════════════════════════════════════════════════╝

type team struct {
	name    string
	captain string
}

type entry struct {
	team
	round int
	score int
}

func main() {
	e := entry{
		name:    "Quizzly Bears",
		captain: "Ada",
		round:   3,
		score:   7,
	}
	fmt.Println(e.name, "captained by", e.captain)

	total := e.score
	if e.score >= 7 {
		total := total + 3
	}
	fmt.Printf("round %d: %d points\n", e.round, total)
}
