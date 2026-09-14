//go:build ignore

package main

import "fmt"

// ╔══════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 003  ·  ch2 constants, fmt  ·  difficulty 2/10  ·  ~3m   ║
// ║ shape: COMPLETE  ·  tags: const-basics, computed-const,           ║
// ║                          printf-verbs                             ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The Red Lion's jukebox plays nonstop from open to close. ║
// ║ The landlord wants to know what it earns on a perfect day.        ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ TASK                                                              ║
// ║  1. Declare constants: the pub is open 11 hours, an hour is 60    ║
// ║     minutes, a song is 3 minutes long, a song costs 0.50.         ║
// ║  2. Declare a constant songsPerDay, calculated from the others    ║
// ║     (no typing 220 by hand).                                      ║
// ║  3. Replace the placeholder with two Printf lines matching the    ║
// ║     output below exactly. Takings are songs × price.              ║
// ║  HINT  Printf needs a verb for a whole number and one for a       ║
// ║        float to 2 decimal places. Don't forget the \n.            ║
// ╠══════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                   ║
// ║  Songs per day: 220                                               ║
// ║  Takings: £110.00                                                 ║
// ╚══════════════════════════════════════════════════════════════════╝

// your constants here
const pubHours int = 11
const minutesInHour int = 60
const songLength int = 3
const songCost float64 = 0.50
const songsPerDay int = pubHours * minutesInHour / songLength
const earningsPerDay float64 = float64(songsPerDay) * songCost


func main() {
	fmt.Printf("Songs per day: %d\nTakings: $%.2f\n", songsPerDay, earningsPerDay)
	
}

// ── VERDICT: pass, notes ─────────────────────────────────────────
// Constant maths correct first go (220, 110.00). Two bounces on output:
// missing trailing \n, then label text. Used `$` where card said `£`:
// card's fault, Matt's in Canada. Future cards use `$`.
// - Typed consts (`const x int = 11`) force the float64() conversion on
//   earningsPerDay. Untyped consts (`const x = 11`) would multiply fine.
// - gofmt -l still flags: double blank line, trailing whitespace in main.
