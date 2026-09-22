//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 015 · ch5 structs · difficulty 2/10 · ~2 min         ║
// ║ shape: PREDICT · tags: struct-zero-values                    ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  A game creates a guest and a named player.          ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                       ║
// ║ 1. Write all three output lines under PREDICTION before     ║
// ║    running the program. Leave the executable code alone.    ║
// ║ 2. Then run it and compare with your prediction.            ║
// ║ HINT: %q prints a string in quotes; %t prints a bool.       ║
// ╚════════════════════════════════════════════════════════════════╝

type player struct {
	name  string
	score int
	ready bool
}

// PREDICTION:
// line 1: guest name is empty, 0 score, and 0 for the bool, which means false
// line 2: name is Mina, 0 score, 0 bool > false
// line 3: guest ready 0 > false, member ready true

func main() {
	var guest player
	fmt.Printf("guest: name=%q score=%d ready=%t\n", guest.name, guest.score, guest.ready)
	member := player{name: "Mina"}
	fmt.Printf("member: name=%q score=%d ready=%t\n", member.name, member.score, member.ready)
	member.ready = true
	fmt.Printf("guest ready=%t member ready=%t\n", guest.ready, member.ready)
}

/*
Verdict: pass, notes. Predicted values correct in prose, not exact output.
Explained %q versus %s and bool zero value: false directly, not numeric conversion.
Matt corrected //line comments to ordinary // line comments before close.
Final gofmt, vet and run pass. No independent mastery inferred from explanation.
*/
