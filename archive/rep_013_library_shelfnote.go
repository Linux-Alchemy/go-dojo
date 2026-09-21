//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 013 · ch3/ch4 · difficulty 3/10 · ~4 min              ║
// ║ shape: PREDICT · tags: guard-clauses, no-ternary               ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The village library prints a note for every book on   ║
// ║ its list. Somebody wants to know what it'll actually say.      ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. Write your predicted output as a comment above main,        ║
// ║    one line per call, BEFORE running anything.                 ║
// ║ 2. Trace every line of shelfNote for each call, top to bottom. ║
// ║ 3. Then go run . and compare. Don't change the code.           ║
// ║                                                                ║
// ║ (Inlay hints off: <leader>uh)                                  ║
// ╚════════════════════════════════════════════════════════════════╝

// shelfNote returns the librarian's note for one book.
func shelfNote(title string, daysOut, loanDays int) string {
	late := daysOut - loanDays
	fine := 1
	if late > 7 {
		fine = 3
	}
	if daysOut == 0 {
		return title + ": on the shelf"
	}
	if late > 0 {
		return fmt.Sprintf("%s: %d days late, fine $%d", title, late, fine)
	}
	return fmt.Sprintf("%s: %d days left", title, -late)
}

// PREDICTION:
// Dune: 11 days left
// Emma: 6 days late, fine $1
// Ulysses: 16 days late, fine $3
// Heidi: on the shelf
// Kim: 0 days left
// Momo: 7 days late, fine $3

func main() {
	fmt.Println(shelfNote("Dune", 3, 14))
	fmt.Println(shelfNote("Emma", 20, 14))
	fmt.Println(shelfNote("Ulysses", 30, 14))
	fmt.Println(shelfNote("Heidi", 0, 14))
	fmt.Println(shelfNote("Kim", 14, 14))
	fmt.Println(shelfNote("Momo", 21, 14))
}

// ── VERDICT: pass, notes ───────────────────────────────────────────
// 5/6 predicted correctly, blind, with the tracing in the right place:
// Heidi hit the daysOut == 0 guard before the late check, and Kim at
// exactly 0 read as "0 days left", not late. Both orderings correct.
// Miss: Momo. late = 21 - 14 = 7, and the override is `late > 7`,
// which is false at exactly 7 — fine stays at its default 1, not 3.
// Same family as rep 004's `< 2` vs `<= 2`: the boundary, not the shape.
// gofmt clean, go vet clean.
