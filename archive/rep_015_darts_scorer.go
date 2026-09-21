//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 015 · ch4 functions · difficulty 3/10 · ~4 min        ║
// ║ shape: PREDICT · tags: closures                                ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  Thursday night darts at the Crown. The scoreboard is  ║
// ║ a Go program, because of course it is.                         ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. Write your predicted output in the PREDICTION comment,      ║
// ║    one line per Println, BEFORE running anything.              ║
// ║ 2. For each line, be able to say which total you're reading     ║
// ║    and what it was a moment before.                            ║
// ║ 3. Then go run . and compare. Don't change the code.           ║
// ║                                                                ║
// ║ (Inlay hints off: <leader>uh)                                  ║
// ╚════════════════════════════════════════════════════════════════╝

// scorer hands back a running-total function for one darts player.
func scorer() func(int) int {
	total := 0
	return func(points int) int {
		total += points
		return total
	}
}

// PREDICTION:
// 20, was 0
// 5 was 0
// 39 was 20
// 6 was 5
// 99 was 39
// 7 was 0
// 9 was 6

func main() {
	alex := scorer()
	sam := scorer()

	fmt.Println(alex(20))
	fmt.Println(sam(5))
	fmt.Println(alex(19))
	fmt.Println(sam(1))
	fmt.Println(alex(60))

	bailey := scorer()
	fmt.Println(bailey(7))
	fmt.Println(sam(3))
}

// ── VERDICT: clean ─────────────────────────────────────────────────
// 7/7 blind, and the "was N" annotations show the tracing was real
// rather than lucky: alex and sam kept separate totals, and bailey
// started at 0 despite being made after 99 points had gone through
// the factory — i.e. `total := 0` runs once per scorer() call, not
// once per program. gofmt clean, go vet clean. No flags.
// closures: revisit → shaky (one clean shape on record, PREDICT).
