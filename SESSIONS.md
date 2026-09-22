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

## 2026-09-20 · 4 reps · asked 2–3 · avg 3.0 · scope: ch1–4 (carried forward, then raised at the close)

**Worked on:** guard-clauses, no-ternary (013), guard-clauses, switch-tagless (014, carded as if-syntax), closures (015), closures, shadowing (016)
**Solid:** closures and shadowing — both on two clean reps of different shapes (PREDICT 015 + FIX 016, and PREDICT 007 + FIX 016). First promotions since the v2.1 migration; both due for retest 2026-09-25.
**Retested:** none due.
**Struggled:** nothing reached not-yet or revisit; guard-clauses and no-ternary both came *off* revisit. Two things still worth naming: 013 missed the `> 7` override at exactly 7 (read the shape right, landed the wrong side of the boundary), and 014 was solved by replacing the if-chain with a tagless switch rather than diagnosing either planted bug, so the boundary evidence there is "wrote it right", not "spotted it wrong".
**Python mind sightings:** none confirmed. Matt offered one — writing tagless cases under a tagged `switch age`, giving `cannot convert age < 16 (untyped bool value) to type int` — and it was talked back out of the column: Python's `match` wouldn't lead there and `if/elif` is a different shape. That's two Go forms crossed, which is a better class of error and worth him knowing the difference.

**Note:** The best session so far, and the shape of it matters more than the count. Closures went from a lesson Matt flagged on the 17th as having jammed him up, to solid in two reps three days later — the PREDICT proved he could trace two closures from one factory keeping separate totals, and the FIX proved he could spot `sold := sold + n` being reborn every call in a program that compiles and vets clean. That is exactly what the lesson-import pipeline was built to do, so keep using it: `lessons/ch4_currying.md` is still sitting there unspent.

No stagnation call-out needed, because Matt made the call himself before I could: three sessions parked at 2–3, and he asked for ch1–4 at 3–4 plus a new ch5 structs band at 2–3, along with an arrangement to report his Boot.dev position as he goes. Exactly the right read. gofmt was also clean on all four reps after being dirty on three earlier ones.

Next session: rep 017 (COMPLETE 3/10, quiz-night `applyRule`) is loaded and unattempted — the first `higher-order` rep, drawn by repair-before-retry because `currying` is at revisit and can't be retried until its prerequisite lands. After that, the new bands open up: ch1–4 at 3–4, and the first structs reps at 2–3 on `struct-define-literal`, `struct-zero-values` and `nested-structs` only. Worth checking the exact-boundary habit again at 4, since it's now been asked three times and answered two different ways.

## 2026-09-21 · 4 reps · asked 2–3 (structs), 3–4 (ch1–4) · avg 2.75 · scope: rep 017 (ch4), then ch5 structs only

**Worked on:** higher-order, guard-clauses (017), struct-define-literal (018, 019), struct-zero-values (019), nested-structs, struct-copy-semantics (020)
**Solid:** none promoted this session (every tag here is on its first shape). Closest: `struct-define-literal`, from revisit to shaky in one rep, and `higher-order`, clean code on a first rep.
**Struggled:** struct-define-literal (named-field literals `ticket{plate: "..."}` were new; tried `ticket{ticket.plate, ...}`, reading a field off the type; needed nudge + snippet, so revisit, recovered on 019). Also missed a planted hours/rate swap on 018 because the total came out the same.
**Python mind sightings:** none confirmed. The `b := a` alias read on 020 is Python's model (assignment binds a name to the same object), but his predictions contradicted it, so it's a wording gap as much as a mental-model one. Worth one pointed rep.
**Note:** The struct reps did what first reps should: they found the gap straight away (named-field literals) and closed most of it within one rep. Two of the four hiccups were my card errors, not his: 017's step 2 was badly worded, and 020 used `++` without explaining it. Card rule added to the ledger. No stagnation call: the bands were raised only yesterday and this session had real struggles, so 2–3 on structs is doing its job. Next session: rep 021 (COMPLETE 3/10, first value-receiver method, boundary at exactly 5 kg) is loaded. After it, a copy-semantics rep where he has to *say* copy or alias, then embedded structs, which he asked for by name and hasn't touched yet.


## 2026-09-22 · 2 reps · asked 2–3 · avg 3.0 · scope: ch5 structs only

**Worked on:** methods-value-receiver (021 COMPLETE, 022 FIX)
**Solid:** no new promotions; closures and shadowing remain solid. Methods moved from revisit after 021 to shaky after 022.
**Struggled:** methods-value-receiver (021 needed multiple nudges on method name versus local variable, receiver field access, surcharge threshold and subtraction). On 022, the code fix was unaided and clean; explanation needed clarification that `lantern{}` creates a new zero-valued lantern and the receiver gets the whole struct copy, not just its fuel number.
**Python mind sightings:** none confirmed this session; receiver and literal terminology needed clarification, without enough evidence to attribute that to Python.
**Note:** Both reps produced the expected output and passed vet; 021 had a redundant `== true` and an extra blank line, while 022 was gofmt-clean. Matt correctly traced the 6-unit lantern through the false guard to 18 minutes, then confirmed the whole-struct-copy explanation made sense. That is progress, but not independent evidence of the copy model yet. Keep structs at 2–3: the band was recently set and these reps are still finding useful gaps; methods are improving, with another independent explanation needed before calling them settled. No difficulty increase suggested.

Paused at Matt's request. Rep 023 (TRANSLATE, anonymous structs, 3/10) remains loaded and unattempted in main.go; resume there at ch5 structs only, 2–3. Reps 021–022 are archived; total completed is 22. No rep was loaded during close.


## 2026-09-19 · recovered 2026-09-22 · 3 historical exercises · avg 2.3 · scope: ch5 typed validation and zero values

**Provenance:** Recovered from the pre-sync stash. These exercises originally used numbers 013–015 in a separate local history; they are identified here by name and stored in [a dated archive](archive/2026-09-19-import/README.md). They do not renumber or increment the current sequence, which remains 22 completed with rep 023 pending. This entry is appended out of date order to preserve existing notes.

**Worked on:** collection validation (COMPLETE 3/10, pass, notes); nested stock (FIX 2/10, clean); struct zero values (PREDICT 2/10, pass, notes). Imported [Boot.dev nested-user validation lesson](lessons/ch5_typed_validation.md).
**Solid:** none promoted. The nested-field FIX supplies one clean shape; other exercises required notes.
**Struggled:** collection validation used a redundant boolean-return branch; zero-value predictions were correct in prose but described bool as numeric zero converting to false. Clarified direct false and `%q` versus `%s`. Prediction comments beginning `//line` caused compiler-directive errors; Matt corrected those before the historical session closed.
**Python mind sightings:** checking possible types of concretely typed fields (reported lesson difficulty); interpreting a bool default through numeric truthiness.
**Note:** Preserve the distinction between Go's type guarantees and the task's permitted values. Historical checks were recorded as passing; they were not rerun as part of recovery. No independent mastery is inferred from the supplied lesson solution or explanations. Newer ratings and dates stay intact; typed-value-validation is restored as shaky, last exercised 2026-09-19. Resume the current pending rep 023, not the old pad.
