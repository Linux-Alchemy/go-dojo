//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 018 · ch5 structs · difficulty 3/10  ·  ~4 min        ║
// ║ shape: FIX · tags: struct-define-literal                       ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The Wharf Street car park prints a slip for every     ║
// ║ car. Tonight's three:                                          ║
// ║   KX19 ABC · bay B12 · 3 hours at $2/h                         ║
// ║   LM70 XYZ · bay C04 · 2 hours at $5/h                         ║
// ║   PQ21 RST · bay A01 · 1 hour  at $4/h                         ║
// ║ It compiles. It runs. The attendant is not happy.              ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. Find out why the slips don't match the cars above.          ║
// ║ 2. Fix the tickets so this mistake can't happen to them again, ║
// ║    even if someone reorders the fields in the struct one day.  ║
// ║ 3. Don't touch the struct or describe.                         ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║   KX19 ABC in bay B12: 3h at $2/h = $6                         ║
// ║   LM70 XYZ in bay C04: 2h at $5/h = $10                        ║
// ║   PQ21 RST in bay A01: 1h at $4/h = $4                         ║
// ╚════════════════════════════════════════════════════════════════╝

type ticket struct {
	plate string
	bay   string
	hours int
	rate  int
}

// describe formats one parking slip.
func describe(t ticket) string {
	return fmt.Sprintf("%s in bay %s: %dh at $%d/h = $%d", t.plate, t.bay, t.hours, t.rate, t.hours*t.rate)
}

func main() {
	a := ticket{plate: "KX19 ABC", bay: "B12", hours: 3, rate: 2}
	b := ticket{plate: "LM70 XYZ", bay: "C04", hours: 2, rate: 5}
	c := ticket{plate: "PQ21 RST", bay: "A01", hours: 1, rate: 4}

	fmt.Println(describe(a))
	fmt.Println(describe(b))
	fmt.Println(describe(c))
}

// ── verdict: pass, notes ─────────────────────────────────────
// Found the plate/bay swap unaided; missed the hours/rate swap on c
// (total $4 either way hid it) until pointed at line 3.
// Named-field literals were new: first tried `ticket{ticket.plate, ...}`
// (field selector on the type). Nudge + snippet (`wheel{radius: 17, ...}`),
// then all three rewritten correctly. gofmt/vet clean.
