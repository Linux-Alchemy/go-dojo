//go:build ignore

package main

import "fmt"

// ╔══════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 004  ·  ch3 conditionals  ·  difficulty 3/10  ·  ~3 min  ║
// ║ shape: FIX  ·  tags: if-syntax, guard-clauses                     ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The corner shop's parcel counter sorts parcels by weight.║
// ║ A customer with a 2kg parcel got charged the medium rate.         ║
// ║ It compiles fine. It's just wrong.                                ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ RULES                                                             ║
// ║  zero or less          → invalid                                  ║
// ║  up to and incl. 2kg   → small                                    ║
// ║  up to and incl. 10kg  → medium                                   ║
// ║  over 10kg             → large                                    ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ TASK                                                              ║
// ║  1. Run it. Compare against the expected output.                  ║
// ║  2. Find the bug in postage and fix it. Leave main alone.         ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                   ║
// ║  0.00kg: invalid                                                  ║
// ║  1.50kg: small                                                    ║
// ║  2.00kg: small                                                    ║
// ║  9.99kg: medium                                                   ║
// ║  10.00kg: medium                                                  ║
// ║  10.50kg: large                                                   ║
// ╚══════════════════════════════════════════════════════════════════╝

// postage returns the parcel size band for a weight in kg.
func postage(weightKg float64) string {
	if weightKg <= 0 {
		return "invalid"
	}
	if weightKg <= 2 {
		return "small"
	}
	if weightKg <= 10 {
		return "medium"
	}
	return "large"
}

func main() {
	fmt.Printf("%.2fkg: %s\n", 0.0, postage(0))
	fmt.Printf("%.2fkg: %s\n", 1.5, postage(1.5))
	fmt.Printf("%.2fkg: %s\n", 2.0, postage(2))
	fmt.Printf("%.2fkg: %s\n", 9.99, postage(9.99))
	fmt.Printf("%.2fkg: %s\n", 10.0, postage(10))
	fmt.Printf("%.2fkg: %s\n", 10.5, postage(10.5))
}

// ── VERDICT: clean ───────────────────────────────────────────────
// Spotted the boundary: `< 2` → `<= 2`, so a 2.00kg parcel is small.
// First run, gofmt and vet clean. No flags.
