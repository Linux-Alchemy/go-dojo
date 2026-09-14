package main

import "fmt"

// ╔══════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 006  ·  ch4 functions  ·  difficulty 3/10  ·  ~4 min     ║
// ║ shape: COMPLETE  ·  tags: multiple-returns, guard-clauses         ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The Red Lion pub quiz pays out in whole dollars. The     ║
// ║ winning team splits the pot evenly; any dollars that won't split  ║
// ║ go in the kitty for next week.                                    ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ TASK                                                              ║
// ║  1. Fill in splitPot. It returns each member's share, then the    ║
// ║     kitty. Leave main alone.                                      ║
// ║  2. Whole dollars only: no floats anywhere.                       ║
// ║  3. Every line of the expected output must match.                 ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                   ║
// ║  $50, 4 members: $12 each, $2 to the kitty                        ║
// ║  $60, 3 members: $20 each, $0 to the kitty                        ║
// ║  $7, 10 members: $0 each, $7 to the kitty                         ║
// ║  $50, 0 members: $0 each, $50 to the kitty                        ║
// ╚══════════════════════════════════════════════════════════════════╝

// splitPot shares a prize pot in whole dollars between team members.
// It returns each member's share and whatever is left for the kitty.
func splitPot(pot, members int) (int, int) {
	return 0, 0
}

func main() {
	share, kitty := splitPot(50, 4)
	fmt.Printf("$50, 4 members: $%d each, $%d to the kitty\n", share, kitty)

	share, kitty = splitPot(60, 3)
	fmt.Printf("$60, 3 members: $%d each, $%d to the kitty\n", share, kitty)

	share, kitty = splitPot(7, 10)
	fmt.Printf("$7, 10 members: $%d each, $%d to the kitty\n", share, kitty)

	share, kitty = splitPot(50, 0)
	fmt.Printf("$50, 0 members: $%d each, $%d to the kitty\n", share, kitty)
}
