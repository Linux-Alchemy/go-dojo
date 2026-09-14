//go:build ignore

package main

import "fmt"

// ╔══════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 001  ·  ch1 variables  ·  difficulty 2/10  ·  ~2 min     ║
// ║ shape: FIX  ·  tags: short-decl                                   ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The car park sign shows free spaces and the hourly rate. ║
// ║ Someone wrote it in a hurry and it doesn't compile.               ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ TASK                                                              ║
// ║  1. Run it. Read the compiler's complaints.                       ║
// ║  2. There are two bugs. Fix both so the output matches below.     ║
// ║  3. Don't change any numbers or the Printf lines.                 ║
// ║  HINT  both bugs are about how a variable gets declared.          ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                   ║
// ║  Free spaces: 33 of 120                                           ║
// ║  Hourly rate: £2.50                                               ║
// ╚══════════════════════════════════════════════════════════════════╝

var spaces int = 120

func main() {
	taken := 87
	rate  := 2.50
	free := spaces - taken
	fmt.Printf("Free spaces: %d of %d\n", free, spaces)
	fmt.Printf("Hourly rate: £%.2f\n", rate)
}

// ── VERDICT: pass, notes ─────────────────────────────────────────
// Output matched. Both bugs found (package-level :=, var + := mashup).
// - gofmt -l flags main.go: `rate  :=` has a stray double space.
// - `var spaces int = 120`: type is redundant, `var spaces = 120`.
// - Two good questions en route on why := can't take a type. Rated shaky
//   as a learning rep, not a miss.
