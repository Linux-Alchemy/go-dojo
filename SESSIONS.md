# Go Dojo — Session Notes

Append-only. One entry per session, written by the go-dojo skill on close. Read the last few before each close for the stagnation check.


## 2026-09-12 · 3 reps · asked 2–3 · avg 2.0 · scope: ch1–4 (warm-up)

**Worked on:** short-decl, int-division, type-conversion, printf-verbs, const-basics, computed-const
**Solid:** none yet (all pass-with-notes; first session, lots of good questions)
**Struggled:** short-decl (tried `var x T := v` and `x T := v`), int-division (expected 24/7 not to truncate), printf-verbs (missed trailing `\n`, label not exact match)
**Python mind sightings:** type-hint declaration `x T := v` ×1, `/` expected to give a float ×1
**Note:** Warm-up after time away. The concepts land once explained; the misses were mostly exactness (newline, label text, gofmt) rather than logic. Also covered gopls inlay hints and the zsh `%` no-newline marker. Opening scope was too hot (strings.Index/slicing), so it was withdrawn and restarted at 2. Next session: resume rep 004 (FIX 3/10), do a blind PREDICT with inlay hints off to firm up int-division, then settle at 3.

## 2026-09-13 · 2 reps · asked 2–3 · avg 3.0 · scope: ch1–4 (low-key)

**Worked on:** if-syntax, guard-clauses, int-division, type-conversion
**Solid:** if-syntax, guard-clauses (spotted the `<` vs `<=` boundary first run)
**Struggled:** none at not-yet/revisit; int-division and type-conversion stay shaky (asked why `float64(45 / 4)` gave 11.00 and why `Println` showed 11 before predicting)
**Python mind sightings:** expected `Println(11.0)` to print `11.0` ×1; converting the result instead of the inputs ×1
**Note:** Short, easy session and it went well: the FIX was clean, and the PREDICT went 6/6 once the two gotchas were explained. Int truncation has now come up in both sessions, so it's a real pattern. Next time: resume rep 006 (COMPLETE 3/10, splitPot), then a truly blind PREDICT where the conversion comes too late, to see whether it's stuck.

## 2026-09-14 · 1 rep · asked 2–3 · avg 3.0 · scope: ch1–4 (carried forward)

**Worked on:** multiple-returns, guard-clauses (rep 006, splitPot)
**Solid:** whole-dollar division and remainder; zero-members guard placed before division. Corrected solution passed all four output cases, gofmt, and go vet.
**Struggled:** multiple-returns (initially returned `pot, 0` in the zero-members branch, following input parameter order instead of the promised share/kitty result order; fixed after one nudge). Both rep tags remain shaky under the one-nudge rating rule.
**Python mind sightings:** no confirmed Python-specific slip; input names and output positions were being conflated.
**Note:** Most of the session went into why return order matters: returned expressions are assigned to the caller's variables by position, independently of parameter order or names. Matt correctly pointed out that `return members, pot` also works inside `if members == 0`; discussed why `return 0, pot` expresses the dollar-share meaning more clearly. Code is correct, but check the positional-return mental model again in a future rep rather than assuming it is settled. Rep 007 (PREDICT 3/10, short-decl + shadowing) is loaded and unattempted; resume there at ch1–4, difficulty 2–3.

## 2026-09-14 · 2 reps · asked 2–3 · avg 3.0 · scope: ch1–4 (carried forward)

**Worked on:** short-decl, shadowing (rep 007), type-conversion, int-division, grouped-params (rep 008)
**Solid:** short-decl, shadowing (blind PREDICT 3/3 first run, understood that the inner `:=` makes a new variable that dies at the brace)
**Struggled:** type-conversion (fixed `float64(total / rounds)` by changing the params to float64 instead of converting the inputs; left a redundant `float64(...)` wrapper). No not-yet or revisit verdicts.
**Python mind sightings:** none confirmed. The signature fix is closer to "make the types fit" than to Python mind.
**Note:** Big step on int division: Matt explained correctly, and without help, why converting after the division is too late. That's the pattern from all three earlier sessions, now understood. The gap has moved to where a conversion belongs: he didn't see that float64 params only compiled because `main` passed bare numbers, which take whatever type is needed, until we talked it through; he then correctly put it as "`:=` locks in int". Still early days at 2–3 (three days), and there are real notes to consolidate, so no push up yet. Next session: rep 009 (COMPLETE 3/10, car park tariff, switch-tagless + multiple-returns), then a rep that passes int *variables* into a function needing a float64 result, to confirm conversion at the use site has stuck.

## 2026-09-16 · 1 rep · asked 2–3 · avg 3.0 · scope: ch1–4 (carried forward) · backfilled from ledger

**Worked on:** switch-tagless, multiple-returns (rep 009, car park tariff)
**Solid:** switch-tagless, multiple-returns (output matched first run, no nudges; boundaries inclusive in the right places; fee-then-band return order right in all four branches)
**Struggled:** none. Notes were cosmetic (gofmt blank line, redundant parens in `(60 * 2)`).
**Python mind sightings:** none.
**Note:** Written after the fact, at the close of the next session, because this session's note was never appended. Most of the session went on finding and removing the stray duplicate pad at `~/learn/go-dojo`. The positional-return worry from 09-14 looks settled.

## 2026-09-16 · 3 reps · asked 2–3 · avg 3.0 · scope: ch1–4 (carried forward)

**Worked on:** type-conversion, int-division (010), no-truthiness, logical-ops (011), no-ternary, guard-clauses, sprintf-vs-printf (012)
**Solid:** type-conversion, int-division (010 converted the int inputs at the use site, signature untouched, first run: the 008 gap is closed); no-truthiness, logical-ops (011 clean, `!score` → `score == 0` and `||` → `&&`)
**Struggled:** guard-clauses, no-ternary, sprintf-vs-printf, all revisit after four nudges on 012 (independent ifs overwriting one result var; an if/else returning on both branches that made the rest of the function unreachable; read `missingStock := qty - inStock` as if it limited the next block to short orders; no-ternary landed as an if/else with a return in each branch, repeating the Sprintf, rather than default-then-override)
**Python mind sightings:** `if !score` (Python's `if not score:`) ×1, caught and fixed unprompted. Also misses the Python ternary, and said so.
**Note:** Two quick wins, then a real wobble, and it wasn't syntax. On 012 the code was written to match what Matt meant rather than what runs: an assignment treated as a condition, and branches expected to stop running once one had "answered". Tracing A101 by hand (3 − 10 = −7, so "short by -7 items") is what made it click, and he found the fix, an outer `qty > inStock` check, himself. Range is unchanged since 09-12, but this session had real struggles, so no push yet. Type conversion and int division, the pattern from every earlier session, are now solid, which is what the reps are for. gofmt was dirty again (format-on-save still broken) and gopls slipped in a stray `golang.org/x/text/message` import. Next session: rep 013 (PREDICT 3/10, library shelfNote, built to force line-by-line tracing past an unconditional subtraction and two boundaries), then a guard-order FIX and a default-then-override COMPLETE to pull the revisit tags back up.
