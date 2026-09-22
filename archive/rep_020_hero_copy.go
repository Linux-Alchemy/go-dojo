//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 020 · ch5 structs · difficulty 3/10  ·  ~4 min        ║
// ║ shape: PREDICT · tags: nested-structs, struct-copy-semantics   ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  A dungeon game's save code. The dev swears copying a  ║
// ║ hero "just makes another name for the same hero". Does it?     ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. Read main top to bottom. Don't run it yet.                  ║
// ║ 2. Write your predicted output as a comment block above main,  ║
// ║    one line per print, exactly as Go would print it.           ║
// ║ 3. Then run it and say check.                                  ║
// ║                                                                ║
// ║ %q prints a string with double quotes around it.               ║
// ║ Inlay hints off (<leader>uh) — no peeking.                     ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║   four lines — yours to predict                                ║
// ╚════════════════════════════════════════════════════════════════╝

type stats struct {
	hp    int
	level int
}

type hero struct {
	name  string
	stats stats
}

func main() {
	a := hero{name: "Wren", stats: stats{hp: 20, level: 3}}
	b := a
	b.name = "Moss"
	b.stats.hp = 5

	fmt.Println(a.name, a.stats.hp)
	fmt.Println(b.name, b.stats.hp, b.stats.level)

	var c hero
	c.stats.level++
	fmt.Printf("%q %d %d\n", c.name, c.stats.hp, c.stats.level)

	a.stats = b.stats
	b.stats.level = 9
	fmt.Println(a.stats.hp, a.stats.level)
}

// Wren, 20
// Moss 5,3
// i don't know what .level++ means
// 5,3

// ── verdict: pass, notes ─────────────────────────────────────
// 3/4 predicted right (lines 1, 2, 4); line 3 skipped — `++` wasn't
// explained on the card (card error: untaught notation). Line 3 is `"" 0 1`.
// Predictions show copy semantics, but the stated conclusion was that the
// dev "didn't lie" (b as another name for a) — output says copy, not alias.
