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
