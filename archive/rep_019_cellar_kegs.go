//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 019 · ch5 structs · difficulty 2/10  ·  ~3 min        ║
// ║ shape: COMPLETE · tags: struct-define-literal,                 ║
// ║                         struct-zero-values                     ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The Crown's cellar book. The landlord wants every keg ║
// ║ on record, including the ones he knows nothing about yet.      ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. Give keg three fields: beer (text), litres (whole number)   ║
// ║    and chilled (yes/no). Names exactly as written; the output  ║
// ║    prints them.                                                ║
// ║ 2. Fill in bitter, stout and mild with named-field literals:   ║
// ║      Bitter · 50 litres · chilled                              ║
// ║      Stout  · 30 litres · not chilled                          ║
// ║      Mild   · just delivered, nothing else known yet           ║
// ║ 3. Don't write out any value Go would fill in by itself.       ║
// ║ 4. Leave empty alone.                                          ║
// ║                                                                ║
// ║ %+v prints a struct with its field names. Already in main.     ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║   {beer:Bitter litres:50 chilled:true}                         ║
// ║   {beer:Stout litres:30 chilled:false}                         ║
// ║   {beer:Mild litres:0 chilled:false}                           ║
// ║   {beer: litres:0 chilled:false}                               ║
// ╚════════════════════════════════════════════════════════════════╝

type keg struct {
	beer string
	litres int
	chilled bool
}

func main() {
	bitter := keg{beer: "Bitter", litres: 50, chilled: true}
	stout := keg{beer: "Stout", litres: 30, chilled: false}
	mild := keg{beer: "Mild"}
	empty := keg{}

	fmt.Printf("%+v\n", bitter)
	fmt.Printf("%+v\n", stout)
	fmt.Printf("%+v\n", mild)
	fmt.Printf("%+v\n", empty)
}

// ── verdict: pass, notes ─────────────────────────────────────
// Output matched 4/4 first run; named-field literals unaided.
// Flags: `chilled: false` written out on stout (step 3 — mild got it
// right); gofmt dirty (struct field alignment).
