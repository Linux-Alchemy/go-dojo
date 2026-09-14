package main

import "fmt"

// splitPot shares a prize pot in whole dollars between team members.
// It returns each member's share and whatever is left for the kitty.
func splitPot(pot, members int) (int, int) {
	if members <= 0 {
		return 0, pot
	}
	share := pot / members
	return share, pot - share*members
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
