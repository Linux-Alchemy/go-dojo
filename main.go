package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 009 · ch3/ch4 · difficulty 3/10 · ~4 min              ║
// ║ shape: COMPLETE · tags: switch-tagless, multiple-returns       ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The car park barrier needs to know what to charge     ║
// ║ and which band to print on the receipt.                        ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. Fill in chargeFor using a tagless switch (no if/else).      ║
// ║ 2. Tariff:                                                     ║
// ║      the first 30 minutes are free      -> $0,  "grace"        ║
// ║      up to 2 hours                      -> $4,  "short"        ║
// ║      up to 8 hours                      -> $12, "day"          ║
// ║      anything longer                    -> $20, "overnight"    ║
// ║ 3. Return the fee first, then the band.                        ║
// ║ 4. Don't touch report or main.                                 ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║  30 min: grace, $0                                             ║
// ║  31 min: short, $4                                             ║
// ║  120 min: short, $4                                            ║
// ║  480 min: day, $12                                             ║
// ║  481 min: overnight, $20                                       ║
// ╚════════════════════════════════════════════════════════════════╝

// chargeFor returns the fee in whole dollars and the tariff band.
func chargeFor(minutes int) (int, string) {
	return 0, ""
}

// report prints one receipt line for a stay.
func report(minutes int) {
	fee, band := chargeFor(minutes)
	fmt.Printf("%d min: %s, $%d\n", minutes, band, fee)
}

func main() {
	report(30)
	report(31)
	report(120)
	report(480)
	report(481)
}
