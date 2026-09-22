# Go Dojo — Ledger

Maintained by the `go-dojo` skill. Hand-edit freely; it's plain markdown.

## Position
- Boot.dev course position: **Chapter 5 — Structs, through Struct Methods** as of 2026-09-21 ("I've gotten through the struct lessons... so we can add struct methods and embedded structs to the challenges now"). Ch5 is ten lessons: Structs in Go · Nested Structs · Anonymous Structs · Embedded Structs · Struct Methods · Memory Layout · Empty Struct · Empty Struct Memory · Update Users · Send Message. He named methods and embedding specifically, so lessons 1–5 are confirmed done; lessons 6–8 (Memory Layout, Empty Struct, Empty Struct Memory) are **not** confirmed and stay held. Ch1–4 finished. Notes exist for all 16 content chapters (17 is the quiz).
- **Position reporting is Matt's call.** He said on 2026-09-20 he'll report where he is in the Boot.dev course as he goes so reps track it, or name a topic outright when he wants one. Take the newest position he gives over anything written here, and don't drill past it. Within a chapter this applies lesson by lesson, not just chapter by chapter: hold `memory-layout` and `empty-struct` until he confirms he's past those two lessons (still held as of 2026-09-21).
- **Current scope (bands set 2026-09-20 session 7; ch5 tag list widened 2026-09-21 on Matt's report):** two bands running together —
  - **ch1–4 at difficulty 3–4** (up from 2–3, which had been carried unchanged since 09-12)
  - **ch5 structs at difficulty 2–3**, new material, starting low
  Structs tags open for selection: `struct-define-literal`, `struct-zero-values`, `nested-structs`, `anonymous-structs`, `embedded-structs`, `methods-value-receiver`, `struct-copy-semantics`. The last four were added 2026-09-21 when Matt reported finishing the struct lessons and asked for methods and embedding by name. `struct-copy-semantics` came with them because the "value receiver can't mutate" tripwire *is* the Struct Methods lesson, and it's ch4 pass-by-value in struct clothing — worth knowing a rep may hit `pass-by-value` (ch4, `unseen`) sideways. Still held: `memory-layout`, `empty-struct`. Difficulty bands unchanged; Matt hasn't asked to move them. He reviewed the original split on 2026-09-20 and called it "accurate enough", so it is agreed, not assumed.
- Previous scope: ch1–4 at 2–3, low-key/low-effort, carried forward through sessions 3–7.
- Reps completed: 20 (session 8: 017–020 all pass-notes; session 7: 013 pass-notes 5/6 blind, 014 pass-notes after one nudge, 015 clean, 016 clean)
- Session 8 (2026-09-21) scope: Matt asked for a few rounds of **ch5 structs only at 2–3** after 017.
- Next: rep 021 in `main.go`, **not yet attempted** (COMPLETE 3/10, Dash & Dump courier `parcel.cost()` value-receiver method; $4 flat + $2/kg over 5 kg, express doubles the whole price; boundary at exactly 5 kg planted; prints commented out so it compiles — two `keeps the compiler quiet` lines to delete once live). Resume here, in the ch5 structs-only 2–3 scope from session 8 unless Matt says otherwise.
- **Lesson imports:** `lessons/` holds Boot.dev lessons Matt has flagged for review (lesson text, his solution, what to drill, rep seeds). A tag rated from a lesson import rather than a rep has `last` = –. Use the lesson file as the model when writing reps for that tag.
- **Pad location:** `~/github/go-dojo`, and only here. A stray duplicate at `~/learn/go-dojo` had been in use by mistake; on 2026-09-16 rep 009's solution and the rep 010 load were copied back here and the duplicate was deleted. That path no longer exists — don't look for it. Project memory lives under the `-home-reaper-github-go-dojo` key; sessions must be started from this directory so the ledger and memory agree.
- Check note: run `gofmt -l main.go`, not `.` — archived reps are kept verbatim with their gofmt warts.
- Setup notes: gopls auto-imported `golang.org/x/text/message` mid-rep 012 (probably from a typo); watch for stray imports. Turn inlay hints off (`<leader>uh`) for PREDICT reps; format-on-save for Go is not firing (gofmt dirty on 003, 009 — confirmed recurring, worth fixing in LazyVim). Money in cards uses `$`.

## Concept Ledger
Ratings: **solid** / **shaky** / **revisit** / **unseen**, on the evidence rules in the skill (`solid` = two clean reps of different shapes). `last` = date last touched by a rep (– for none, or a lesson import). `retest` = date a `solid` tag is next due. `needs` = prerequisite tag, where one genuinely exists. Selection: overdue retests, then repair-before-retry, then random within scope weighted by rating (revisit 2×, shaky 1.5×, solid 0.5×).

| Tag | Chapter | Rating | last | retest | needs | Notes |
|---|---|---|---|---|---|---|
| short-decl (`:=` vs `=` vs `var`) | 1 | shaky | 2026-09-14 | – | – | asked why `var x T :=` and `x T :=` fail; python type-hint reflex. 007: inner `:=` new var read correctly. 2026-09-17 (v2.1 migration): solid → shaky; one clean shape on record (PREDICT 007), needs a second, different shape |
| zero-values | 1 | unseen | – | – | – | |
| type-conversion (int↔float, truncation) | 1 | shaky | 2026-09-16 | – | – | asked why float64(a/b) gave 11.00; convert inputs, not result. 008: explained the bug correctly unprompted, but fixed it by making the params float64. 010: converted int inputs at the use site, signature untouched, first run. 2026-09-17 (v2.1 migration): solid → shaky; one clean shape on record (TRANSLATE 010), needs a second, different shape |
| int-division | 1 | shaky | 2026-09-16 | – | – | predicted right after explanation; Println drops .0 on floats was news. 008: explained truncation-before-conversion correctly. 010: `float64(minutes) / 60` first run, no nudge. 2026-09-17 (v2.1 migration): solid → shaky; one clean shape on record (TRANSLATE 010), needs a second, different shape |
| unused-vars-imports | 1 | unseen | – | – | – | |
| shadowing | 1 | solid | 2026-09-20 | 2026-09-25 | – | blind PREDICT 3/3 first run (007). 2026-09-17 (v2.1 migration): solid → shaky pending a second shape. 016: clean FIX, found the `:=` inside the closure unaided — **solid** on two clean shapes (PREDICT 007 + FIX 016) |
| strings-bytes (len, indexing) | 1 | unseen | – | – | – | |
| const-basics | 2 | shaky | 2026-09-12 | – | – | typed every const (`const x int = 11`), forcing conversions |
| computed-const | 2 | shaky | 2026-09-12 | – | – | maths right first go; bounced on output format |
| untyped-const | 2 | unseen | – | – | – | |
| printf-verbs (`%v %d %s %f %T %t %q`) | 2 | shaky | 2026-09-12 | – | – | verbs right; missed trailing \n, label not exact |
| sprintf-vs-printf | 2 | revisit | 2026-09-16 | – | – | Sprintf used right once guided; first draft assigned `fmt.Println(...)` to a string |
| string-concat-types | 2 | unseen | – | – | – | |
| if-syntax (braces, else placement) | 3 | shaky | 2026-09-13 | – | – | clean first run. 2026-09-17 (v2.1 migration): solid → shaky; one clean shape on record (FIX 004), needs a second, different shape |
| no-truthiness | 3 | shaky | 2026-09-16 | – | – | fixed `if !score` → `score == 0` first run. 2026-09-17 (v2.1 migration): solid → shaky; one clean shape on record (FIX 011), needs a second, different shape |
| logical-ops (and, or, not) | 3 | shaky | 2026-09-16 | – | – | spotted `\|\|` should be `&&` for joker rule first run. 2026-09-17 (v2.1 migration): solid → shaky; one clean shape on record (FIX 011), needs a second, different shape |
| if-init | 3 | unseen | – | – | – | |
| switch-value | 3 | unseen | – | – | – | |
| switch-tagless | 3 | shaky | 2026-09-20 | – | – | four-branch tariff, correct fall-through order and inclusive boundaries first run. 2026-09-17 (v2.1 migration): solid → shaky; one clean shape on record (COMPLETE 009), needs a second, different shape. 014: reached for tagless unprompted and ordered the cases right, but wrote the cases under a tagged `switch age` — one nudge, so still not a clean second shape |
| no-ternary | 3 | shaky | 2026-09-20 | – | – | surprised Go has none (likes Python's). 012: used an if/else with a return in each branch, repeating the Sprintf, instead of default-then-override. 013: read the default-then-override shape correctly but applied the override at the wrong boundary (`late > 7` is false at 7) — revisit → shaky |
| func-signature (type-after-name) | 4 | unseen | – | – | – | |
| grouped-params | 4 | shaky | 2026-09-14 | – | – | changed `total, rounds int` to float64; fine syntax, questionable API call |
| multiple-returns | 4 | shaky | 2026-09-16 | – | – | 006: mixed parameter order with result order, one nudge. 009 recheck: fee-then-band correct in all four branches, first run, no nudge — the positional model has landed. 2026-09-17 (v2.1 migration): solid → shaky; one clean shape on record (COMPLETE 009), needs a second, different shape |
| blank-identifier | 4 | unseen | – | – | – | |
| named-returns | 4 | unseen | – | – | – | |
| naked-return | 4 | unseen | – | – | – | |
| pass-by-value | 4 | unseen | – | – | – | |
| guard-clauses | 4 | shaky | 2026-09-21 | – | – | 006: guard placed correctly; returned values reversed, fixed after one nudge. 012: stacked ifs overwriting a result var; if/else returning on both sides made later code unreachable; read `d := a - b` as a condition. Four nudges. 013: traced guard *order* correctly blind (zero-guard before the late check; exactly 0 read as not-late) — revisit → shaky. 014: both boundaries right (16 → $12, 65 → $9), but reached by rewriting the chain as a switch rather than diagnosing the planted bugs. 017: `a < 0 \|\| b < 0` guard, exact boundary at 0 right (0,7 → 7) — the 013 miss not repeated |
| func-as-value (func types) | 4 | unseen | – | – | – | |
| no-defaults-no-kwargs | 4 | unseen | – | – | – | |
| higher-order (func params, func types) | 4 | shaky | 2026-09-21 | – | – | currying leans on it: drill this alone before any currying rep (first two seeds in `lessons/ch4_currying.md`). 017: COMPLETE, `return rule(a, b)` behind the guard, first run — pass, notes only because the reworded card told him where the guard went. One shape on record; needs a second |
| anonymous-funcs | 4 | unseen | – | – | – | |
| currying (func returning func) | 4 | revisit | – | – | higher-order | no rep yet. 2026-09-17: passed the Boot.dev `getLogger` assignment but found it quite a bit harder than closures. Lesson text, his solution, rep ladder and what to pre-supply in `lessons/ch4_currying.md` |
| closures (capture by reference) | 4 | solid | 2026-09-20 | 2026-09-25 | – | 2026-09-17: passed the Boot.dev `adder()` assignment but Matt flagged the lesson as one that jammed him up; lesson text, his solution and rep seeds in `lessons/ch4_closures.md`. 015: 7/7 blind PREDICT, two closures from one factory kept separate totals and a third made late started at 0 — revisit → shaky. 016: clean FIX, spotted `sold := sold + n` reborn every call while the enclosed one never moved — **solid** on two clean shapes |
| defer (LIFO, args evaluated early) | 4 | unseen | – | – | – | |
| struct-define-literal (named vs positional) | 5 | shaky | 2026-09-21 | – | – | 018: found the plate/bay swap, missed hours/rate (same total hid it). Named-field literals new to him — Boot.dev only shows them in passing (anonymous structs). Tried `ticket{ticket.plate, ...}`; needed nudge + snippet → revisit. Matt flagged it himself as one to revisit. 019: COMPLETE, named-field literals right first run, no nudges → shaky. Needs a clean rep of another shape |
| struct-zero-values | 5 | shaky | 2026-09-21 | – | – | 019: omitted everything on the mild keg correctly, but wrote `chilled: false` out on the stout (asked not to write zero values) |
| nested-structs | 5 | shaky | 2026-09-21 | – | – | 020: PREDICT, read `a.stats.hp` / `b.stats.level` through the nested struct correctly on every line he attempted |
| anonymous-structs | 5 | unseen | – | – | – | opened for selection 2026-09-21 (lesson passed on Boot.dev, no rep yet). Seeds S13 |
| embedded-structs (promotion, literal keys) | 5 | unseen | – | – | – | opened for selection 2026-09-21, asked for by name. No rep yet. Two tripwires to plant: literal keys use the **type name** (`sender{user: user{...}}`, not `sender{name: ...}`), and embedding is not inheritance — a func taking a `user` won't take a `sender`. Seeds S7, S8, S15 |
| methods-value-receiver | 5 | unseen | – | – | – | opened for selection 2026-09-21, asked for by name. No rep yet. Receiver is `(r rect)`, never `self`/`this` — a live python-mind risk. Value receiver compiles, runs, and mutates nothing. Seeds S9, S10, S14 |
| struct-copy-semantics | 5 | shaky | 2026-09-21 | – | – | opened 2026-09-21; same rule as ch4 `pass-by-value`. 020: PREDICT — values right (a unchanged after `b := a` edits, nested stats copied), but his stated conclusion was that `b := a` is "another name for the same hero" (alias). Predictions and model disagree; next rep on this tag should make him say which, not just predict. Seeds S10, S11 |
| memory-layout (field order, padding) | 5 | unseen | – | – | – | |
| empty-struct | 5 | unseen | – | – | – | |
| interface-declare-satisfy | 6 | unseen | – | – | – | |
| implicit-implementation (use-site errors) | 6 | unseen | – | – | – | |
| interface-only-methods (no field access) | 6 | unseen | – | – | – | |
| multiple-interfaces | 6 | unseen | – | – | – | |
| type-assertion (ok form, panic form) | 6 | unseen | – | – | – | |
| type-switch | 6 | unseen | – | – | – | |
| empty-interface-any | 6 | unseen | – | – | – | |
| clean-interfaces (small, unaware of types) | 6 | unseen | – | – | – | |
| error-check-return (err != nil, zero values) | 7 | unseen | – | – | – | |
| errors-new-errorf | 7 | unseen | – | – | – | |
| custom-error-type (Error() method) | 7 | unseen | – | – | – | |
| err-shadowing-in-chains | 7 | unseen | – | – | – | |
| panic-recover (when not to) | 7 | unseen | – | – | – | |
| parse-input-errors (strconv) | 7 | unseen | – | – | – | |
| for-three-clause | 8 | unseen | – | – | – | |
| for-as-while (omitted clauses) | 8 | unseen | – | – | – | |
| range-int (1.22) | 8 | unseen | – | – | – | |
| modulo | 8 | unseen | – | – | – | |
| continue-break (innermost only) | 8 | unseen | – | – | – | |
| nested-loops-bounds | 8 | unseen | – | – | – | |
| loop-int-float-mix | 8 | unseen | – | – | – | |
| arrays-vs-slices (fixed type, [:] conversion) | 9 | unseen | – | – | – | |
| slicing-bounds (low:high) | 9 | unseen | – | – | – | |
| make-len-cap | 9 | unseen | – | – | – | |
| append-reassign | 9 | unseen | – | – | – | |
| variadic-spread | 9 | unseen | – | – | – | |
| range-slice (index, copy of element) | 9 | unseen | – | – | – | |
| range-string-runes | 9 | unseen | – | – | – | |
| slice-of-slices | 9 | unseen | – | – | – | |
| shared-backing-array (tricky slices) | 9 | unseen | – | – | – | |
| index-out-of-range (runtime panic) | 9 | unseen | – | – | – | |
| nil-vs-empty-slice | 9 | unseen | – | – | – | |
| map-make-literal (nil map write panic) | 10 | unseen | – | – | – | |
| map-zero-value-read | 10 | unseen | – | – | – | |
| comma-ok | 10 | unseen | – | – | – | |
| map-delete-mutate-by-ref | 10 | unseen | – | – | – | |
| map-counting-idiom | 10 | unseen | – | – | – | |
| map-key-types (comparable, struct keys) | 10 | unseen | – | – | – | |
| nested-maps (inner init) | 10 | unseen | – | – | – | |
| map-iteration-order | 10 | unseen | – | – | – | |
| map-as-set | 10 | unseen | – | – | – | |
| pointer-syntax (&, *, *T) | 11 | unseen | – | – | – | |
| write-through-pointer | 11 | unseen | – | – | – | |
| nil-pointer-guard | 11 | unseen | – | – | – | |
| pointer-receivers (mutate, consistency) | 11 | unseen | – | – | – | |
| method-sets-and-interfaces (*T vs T) | 11 | unseen | – | – | – | |
| pointer-passed-by-value (reassign vs write) | 11 | unseen | – | – | – | |
| pointers-in-maps-structs | 11 | unseen | – | – | – | |
| package-main-vs-library | 12 | unseen | – | – | – | |
| exported-capitalisation | 12 | unseen | – | – | – | |
| package-per-directory | 12 | unseen | – | – | – | |
| go-run-build-install | 12 | unseen | – | – | – | |
| stale-binary | 12 | unseen | – | – | – | |
| go-mod-replace-work | 12 | unseen | – | – | – | |
| import-paths (module + subdir) | 12 | unseen | – | – | – | |
| go-get-tidy | 12 | unseen | – | – | – | |
| clean-package-api | 12 | unseen | – | – | – | |
| goroutine-go-keyword | 13 | unseen | – | – | – | |
| channel-send-receive (blocking) | 13 | unseen | – | – | – | |
| deadlock-diagnosis | 13 | unseen | – | – | – | |
| token-channel (struct{}) | 13 | unseen | – | – | – | |
| buffered-channels | 13 | unseen | – | – | – | |
| close-and-ok | 13 | unseen | – | – | – | |
| range-over-channel | 13 | unseen | – | – | – | |
| select (ok checks) | 13 | unseen | – | – | – | |
| select-default-nonblocking | 13 | unseen | – | – | – | |
| channel-direction-types | 13 | unseen | – | – | – | |
| mutex-lock-defer-unlock | 14 | unseen | – | – | – | |
| race-condition-recognition | 14 | unseen | – | – | – | |
| rwmutex-readers-writers | 14 | unseen | – | – | – | |
| mutex-copy-and-reentrancy | 14 | unseen | – | – | – | |
| race-detector | 14 | unseen | – | – | – | |
| type-parameters-syntax | 15 | unseen | – | – | – | |
| zero-value-of-T | 15 | unseen | – | – | – | |
| constraints-any-vs-methods | 15 | unseen | – | – | – | |
| type-lists-ordered-tilde | 15 | unseen | – | – | – | |
| parametric-interfaces | 15 | unseen | – | – | – | |
| generic-types-methods | 15 | unseen | – | – | – | |
| type-inference-explicit | 15 | unseen | – | – | – | |
| type-definitions (distinct types, conversion) | 16 | unseen | – | – | – | |
| iota-basics (repeat rule, reset per block) | 16 | unseen | – | – | – | |
| iota-skip-offset (_ first, shifts) | 16 | unseen | – | – | – | |
| enum-string-method | 16 | unseen | – | – | – | |
| enum-validation | 16 | unseen | – | – | – | |

## Struggle Patterns
- (durable patterns only — "python mind" specifics go here once observed)
- **int division / conversion timing** (seen 09-12, 09-13, 09-14): expects `/` to keep the fraction; `float64(a / b)` looked like it should rescue it. Keep seeding PREDICTs where the conversion comes too late. 09-14: explained the truncation-before-conversion bug correctly without help, so the model looks to be landing.
- Watch: fixes a type mismatch by changing the function signature (int params → float64) instead of converting inside the function, seen 2026-09-14. Understood why after talking about constants (no type until used) versus variables (type fixed when created).
- **Control flow as intent, not execution** (seen 2026-09-16, rep 012): wrote `missingStock := qty - inStock` expecting the following block to run only when the order was short. Also stacked independent ifs that each overwrote one result, and an if/else returning on both sides, which made the rest of the function unreachable. Tracing A101 by hand landed it. Seed PREDICTs that make him trace every line, plus guard-order FIXes.
- Watch: redundant parens around whole expressions (`(60 * 2)` in 009, `(float64(minutes) / 60)` and the return in 010). Cosmetic, but seen two reps running.
- Card rule, learned 2026-09-21: `x++` / `x--` haven't been taught (they arrive with loops, ch8). Don't use them on a card without a one-line explanation.
- Watch: python type-hint reflex on declarations (`x T := v`), seen 2026-09-12, not yet durable.
- Watch: expects `Println` of a whole float64 to show `11.0` (Python repr), seen 2026-09-13.
- **Watch: the boundary at the exact value.** Reads a comparison's shape correctly and then lands the wrong side of `>` vs `>=`. Caught it unaided on 004 (`< 2` → `<= 2`) and on 009's tariff; missed it on 013 (`late > 7` is false at exactly 7). Not yet durable — it goes both ways — but it is the same question three times. Keep planting the exact-boundary case in expected output so a wrong call shows.
- **Watch: right answer, wrong story** (020, 2026-09-21): predicted `b := a` copy behaviour correctly on every line, then concluded the dev was right that it's "another name for the same hero". Next copy-semantics rep should ask him to state copy vs alias in words, not just predict.
- Watch: on a FIX, rewrites the construct rather than diagnosing the planted bug (014: swapped the whole if-chain for a tagless switch, which was better code but meant neither planted bug was ever named). Good instinct, weaker evidence. Seen once; if it recurs, ask him to name the bug before he fixes it.
- Resolved-ish: gofmt came back clean on all four reps this session (013–016), after being dirty on 003, 009 and 012. Either format-on-save got fixed or he's running it by hand. Stop treating it as a standing issue unless it reappears.

## Rep History
| Rep | Shape | Diff | Tags | Verdict | Time | Note |
|---|---|---|---|---|---|---|
| 001 | FIX | 2 | short-decl | pass, notes | – | gofmt double space; redundant `int` on package var |
| 002 | PREDICT | 2 | int-division, type-conversion, printf-verbs | pass, notes | – | ran before predicting; int/int truncation was news |
| 003 | COMPLETE | 2 | const-basics, computed-const, printf-verbs | pass, notes | – | 2 bounces: no trailing \n, wrong label; gofmt dirty (`£` on card was a card error) |
| 004 | FIX | 3 | if-syntax, guard-clauses | clean | ~3 | `< 2` → `<= 2`, first run |
| 005 | PREDICT | 3 | int-division, type-conversion | pass, notes | – | 6/6 but asked about float64(a/b) and Println(11.0) first |
| 006 | COMPLETE | 3 | multiple-returns, guard-clauses | pass, notes | – | zero-members return order reversed; fixed after one nudge; gofmt/vet clean |
| 007 | PREDICT | 3 | short-decl, shadowing | clean | ~3 | 3/3 blind; inner `:=` scoped to the if block |
| 008 | FIX | 3 | type-conversion, int-division, grouped-params | pass, notes | – | explained the bug right; fixed by making params float64, leaving a redundant float64(...) |
| 009 | COMPLETE | 3 | switch-tagless, multiple-returns | pass, notes | – | output matched first run, no nudges; notes were cosmetic only (gofmt blank line, redundant parens in `(60 * 2)`) |
| 010 | TRANSLATE | 3 | type-conversion, int-division | pass, notes | – | converted inputs not result, signature kept, first run; only flag was redundant parens around both expressions |
| 011 | FIX | 3 | no-truthiness, logical-ops | clean | – | `!score` → `score == 0`, `\|\|` → `&&`; Matt: "really just about choosing the right comparison operators" |
| 012 | COMPLETE | 3 | no-ternary, guard-clauses, sprintf-vs-printf | pass, notes | >8 | matched after four nudges (overwrites, unreachable, assignment-as-check, nesting); gofmt dirty; else-after-return; zero guard second |
| 013 | PREDICT | 3 | guard-clauses, no-ternary | pass, notes | – | 5/6 blind; guard order traced right on both boundary cases; missed the `> 7` override at exactly 7 (Momo, fine $1 not $3); gofmt/vet clean |
| 014 | FIX | 3 | guard-clauses, if-syntax (→ switch-tagless in practice) | pass, notes | – | output matched all 7; replaced the if-chain with a tagless switch instead of diagnosing the two planted bugs; one nudge — tagless cases under a tagged `switch age` (`untyped bool value` → int); gofmt/vet clean |
| 015 | PREDICT | 3 | closures | clean | – | 7/7 blind; annotated each line with the prior total, so the tracing was real; separate factories, separate totals, late closure starts at 0; gofmt/vet clean |
| 016 | FIX | 3 | closures, shadowing | clean | – | `sold := sold + n` → `sold += n` first run, no nudges; both tags promoted to solid; gofmt/vet clean |
| 017 | COMPLETE | 3 | higher-order, guard-clauses | pass, notes | – | 5/5 first run; card step 2 unclear, reworded mid-rep (the reword named the guard placement); boundary at 0 right; gofmt/vet clean |
| 018 | FIX | 3 | struct-define-literal | pass, notes | – | positional → named fields; missed the hours/rate swap until pointed at; `ticket.plate` on the type; nudge + snippet; gofmt/vet clean |
| 019 | COMPLETE | 2 | struct-define-literal, struct-zero-values | pass, notes | – | 4/4 first run, named literals unaided; `chilled: false` written on stout; gofmt dirty (field alignment) |
| 020 | PREDICT | 3 | nested-structs, struct-copy-semantics | pass, notes | – | 3/4 right; line 3 skipped — `++` not on the card (card error). Called `b := a` an alias while predicting copy behaviour |

## Session Log
- 2026-09-09: dojo created; rep 001 loaded.
- 2026-09-12 (session 1): reps 001–003 at 2/10, all pass-with-notes; rep 004 loaded, paused for the night.
- 2026-09-13 (session 2): rep 004 clean; 005 pass-notes; rep 006 loaded, paused before attempt.
- 2026-09-14 (session 3): rep 006 pass-notes at 3/10 after one nudge; discussed parameter order versus positional return values and why `members, pot` works in the zero-members guard. Rep 007 loaded, paused before attempt; scope stays ch1–4 at 2–3.
- 2026-09-14 (session 4): rep 007 clean (blind PREDICT, shadowing); 008 pass-notes (fixed int division by changing the signature). Rep 009 loaded, paused before attempt. Talked about untyped constants versus typed variables (why float64 params only worked with literal arguments). Scope stays ch1–4 at 2–3.
- 2026-09-16 (session 5): rep 009 pass-notes at 3/10 — correct first run, no nudges, both tags promoted to solid (cosmetic flags only). Discovered a stray duplicate pad at `~/learn/go-dojo` that Matt had been working in by mistake; rep 009 and the rep 010 load were copied back into `~/github/go-dojo` and the duplicate removed. Rep 010 loaded, paused before attempt. Scope stays ch1–4 at 2–3.
- 2026-09-16 (session 6): 010 pass-notes (conversion at the use site, first run: type-conversion and int-division now solid); 011 clean (no-truthiness, logical-ops); 012 pass-notes after four nudges (control flow: assignment read as a condition, both-branch returns made later code unreachable), so guard-clauses, no-ternary and sprintf go to revisit. Talked about Go having no ternary. Rep 013 (PREDICT, trace drill) loaded, paused before attempt. Scope stays ch1–4 at 2–3. Session 5's SESSIONS.md entry had been missed and was backfilled.
- 2026-09-17 (lesson import, no reps): Matt brought in Boot.dev lessons that jammed him up. Closures added as `lessons/ch4_closures.md` and currying as `lessons/ch4_currying.md`; both tags set to revisit, currying reported as the harder of the two. Chapter 4 now finished on Boot.dev. Rep 013 still loaded and unattempted.
- 2026-09-17 (ledger migration, no reps): skill moved to v2.1.0. Ledger gained `retest` and `needs` columns; `last` converted from rep numbers to dates. All nine `solid` tags dropped to `shaky` under the two-clean-shapes rule (each has one clean shape on record, named in its Notes), so nothing is due for retest yet. `needs` set for currying → higher-order only. Close phrase is `dojo close`.
- 2026-09-20 (session 7): four reps, all at 3/10 — 013 pass-notes (5/6 blind; missed the `> 7` override at exactly 7), 014 pass-notes (matched, but solved by replacing the if-chain with a tagless switch after one nudge on the switch tag), 015 clean and 016 clean, which took **closures** and **shadowing** to solid — the first promotions since the v2.1 migration, and closures was a lesson Matt had flagged as jamming him up three days earlier. Talked through tagged vs tagless switch (`switch true`) and agreed it wasn't python mind, just two Go forms crossed. Matt then reported Boot.dev position (ch5 structs, about halfway) and **raised the scope himself**: ch1–4 to 3–4, ch5 structs at 2–3, with the back half of ch5 held back by agreement. Rep 017 (COMPLETE 3/10, higher-order) loaded and unattempted.
- 2026-09-21 (scope update, no reps): Matt reported finishing the ch5 struct lessons on Boot.dev and asked for **struct methods** and **embedded structs** in the mix. `anonymous-structs`, `embedded-structs`, `methods-value-receiver` and `struct-copy-semantics` opened for selection; `memory-layout` and `empty-struct` still held, since he named methods and embedding but not the memory lessons. Difficulty bands untouched (ch1–4 at 3–4, ch5 at 2–3). Rep 017 left in the pad, still unattempted.
- 2026-09-21 (session 8): four reps, all pass-notes, avg 2.75. 017 (higher-order) matched first run after a card reword. Then structs only at 2–3 by request: 018 exposed named-field literals as new (Boot.dev only shows them in passing) — nudge + snippet, revisit; 019 wrote them unaided → shaky; 020 predicted copy semantics right but described `b := a` as an alias, and skipped a line because `++` wasn't on the card (card error). Rep 021 (first method) loaded, unattempted. Paused by Matt.
- 2026-09-12: original rep 001 (splitName) withdrawn before attempt — relied on slicing (ch9) and strings.Index Matt hasn't met. Replaced with a ch1 FIX at 2/10.
