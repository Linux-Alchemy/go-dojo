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
//

func main() {
	fmt.Println(shelfNote("Dune", 3, 14))
	fmt.Println(shelfNote("Emma", 20, 14))
	fmt.Println(shelfNote("Ulysses", 30, 14))
	fmt.Println(shelfNote("Heidi", 0, 14))
	fmt.Println(shelfNote("Kim", 14, 14))
	fmt.Println(shelfNote("Momo", 21, 14))
}
