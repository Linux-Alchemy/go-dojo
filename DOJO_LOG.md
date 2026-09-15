# Go Dojo — Ledger

Maintained by the `go-dojo` skill. Hand-edit freely; it's plain markdown.

## Position
- Boot.dev course position: Chapter 4 — Functions, on closures (lesson 16 of 17; currying last) as of 2026-09-09. Notes exist for all 16 content chapters (17 is the quiz). Scope for reps is whatever Matt names; don't drill a chapter he hasn't reached on Boot.dev unless asked.
- Last scope asked: ch1–4, difficulty 2–3, low-key/low-effort — carried forward for sessions 3 and 4 on 2026-09-14
- Reps completed: 8 (session 4: 007 clean, 008 pass-notes)
- Next: rep 009 in `main.go`, **not yet attempted** (COMPLETE 3/10, car park tariff: switch-tagless + multiple-returns). Session 4 (2026-09-14): carried scope ch1–4 at 2–3.
- Check note: run `gofmt -l main.go`, not `.` — archived reps are kept verbatim with their gofmt warts.
- Setup notes: turn inlay hints off (`<leader>uh`) for PREDICT reps; format-on-save for Go looked like it wasn't firing (gofmt dirty twice), check if it recurs. Money in cards uses `$`.

## Concept Ledger
Ratings: **solid** / **shaky** / **revisit** / **unseen**. `last` = rep number when last touched. Selection is random within Matt's scope; ratings only weight the roll (revisit 2×, shaky 1.5×, solid 0.5×).

| Tag | Chapter | Rating | last | Notes |
|---|---|---|---|---|
| short-decl (`:=` vs `=` vs `var`) | 1 | solid | 007 | asked why `var x T :=` and `x T :=` fail; python type-hint reflex. 007: inner `:=` new var read correctly |
| zero-values | 1 | unseen | – | |
| type-conversion (int↔float, truncation) | 1 | shaky | 008 | asked why float64(a/b) gave 11.00; convert inputs, not result. 008: explained the bug correctly unprompted, but fixed it by making the params float64 |
| int-division | 1 | shaky | 008 | predicted right after explanation; Println drops .0 on floats was news. 008: explained truncation-before-conversion correctly |
| unused-vars-imports | 1 | unseen | – | |
| shadowing | 1 | solid | 007 | blind PREDICT 3/3 first run |
| strings-bytes (len, indexing) | 1 | unseen | – | |
| const-basics | 2 | shaky | 003 | typed every const (`const x int = 11`), forcing conversions |
| computed-const | 2 | shaky | 003 | maths right first go; bounced on output format |
| untyped-const | 2 | unseen | – | |
| printf-verbs (`%v %d %s %f %T %t %q`) | 2 | shaky | 003 | verbs right; missed trailing \n, label not exact |
| sprintf-vs-printf | 2 | unseen | – | |
| string-concat-types | 2 | unseen | – | |
| if-syntax (braces, else placement) | 3 | solid | 004 | clean first run |
| no-truthiness | 3 | unseen | – | |
| logical-ops (and, or, not) | 3 | unseen | – | |
| if-init | 3 | unseen | – | |
| switch-value | 3 | unseen | – | |
| switch-tagless | 3 | unseen | – | |
| no-ternary | 3 | unseen | – | |
| func-signature (type-after-name) | 4 | unseen | – | |
| grouped-params | 4 | shaky | 008 | changed `total, rounds int` to float64; fine syntax, questionable API call |
| multiple-returns | 4 | shaky | 006 | mixed parameter order with result order; fixed after one nudge, then discussed positional assignment; correctly identified members equals 0 in the guard; recheck understanding later |
| blank-identifier | 4 | unseen | – | |
| named-returns | 4 | unseen | – | |
| naked-return | 4 | unseen | – | |
| pass-by-value | 4 | unseen | – | |
| guard-clauses | 4 | shaky | 006 | guard placed correctly; returned values reversed, fixed after one nudge |
| func-as-value (func types) | 4 | unseen | – | |
| no-defaults-no-kwargs | 4 | unseen | – | |
| higher-order (func params, func types) | 4 | unseen | – | |
| anonymous-funcs | 4 | unseen | – | |
| currying (func returning func) | 4 | unseen | – | |
| closures (capture by reference) | 4 | unseen | – | |
| defer (LIFO, args evaluated early) | 4 | unseen | – | |
| struct-define-literal (named vs positional) | 5 | unseen | – | |
| struct-zero-values | 5 | unseen | – | |
| nested-structs | 5 | unseen | – | |
| anonymous-structs | 5 | unseen | – | |
| embedded-structs (promotion, literal keys) | 5 | unseen | – | |
| methods-value-receiver | 5 | unseen | – | |
| struct-copy-semantics | 5 | unseen | – | |
| memory-layout (field order, padding) | 5 | unseen | – | |
| empty-struct | 5 | unseen | – | |
| interface-declare-satisfy | 6 | unseen | – | |
| implicit-implementation (use-site errors) | 6 | unseen | – | |
| interface-only-methods (no field access) | 6 | unseen | – | |
| multiple-interfaces | 6 | unseen | – | |
| type-assertion (ok form, panic form) | 6 | unseen | – | |
| type-switch | 6 | unseen | – | |
| empty-interface-any | 6 | unseen | – | |
| clean-interfaces (small, unaware of types) | 6 | unseen | – | |
| error-check-return (err != nil, zero values) | 7 | unseen | – | |
| errors-new-errorf | 7 | unseen | – | |
| custom-error-type (Error() method) | 7 | unseen | – | |
| err-shadowing-in-chains | 7 | unseen | – | |
| panic-recover (when not to) | 7 | unseen | – | |
| parse-input-errors (strconv) | 7 | unseen | – | |
| for-three-clause | 8 | unseen | – | |
| for-as-while (omitted clauses) | 8 | unseen | – | |
| range-int (1.22) | 8 | unseen | – | |
| modulo | 8 | unseen | – | |
| continue-break (innermost only) | 8 | unseen | – | |
| nested-loops-bounds | 8 | unseen | – | |
| loop-int-float-mix | 8 | unseen | – | |
| arrays-vs-slices (fixed type, [:] conversion) | 9 | unseen | – | |
| slicing-bounds (low:high) | 9 | unseen | – | |
| make-len-cap | 9 | unseen | – | |
| append-reassign | 9 | unseen | – | |
| variadic-spread | 9 | unseen | – | |
| range-slice (index, copy of element) | 9 | unseen | – | |
| range-string-runes | 9 | unseen | – | |
| slice-of-slices | 9 | unseen | – | |
| shared-backing-array (tricky slices) | 9 | unseen | – | |
| index-out-of-range (runtime panic) | 9 | unseen | – | |
| nil-vs-empty-slice | 9 | unseen | – | |
| map-make-literal (nil map write panic) | 10 | unseen | – | |
| map-zero-value-read | 10 | unseen | – | |
| comma-ok | 10 | unseen | – | |
| map-delete-mutate-by-ref | 10 | unseen | – | |
| map-counting-idiom | 10 | unseen | – | |
| map-key-types (comparable, struct keys) | 10 | unseen | – | |
| nested-maps (inner init) | 10 | unseen | – | |
| map-iteration-order | 10 | unseen | – | |
| map-as-set | 10 | unseen | – | |
| pointer-syntax (&, *, *T) | 11 | unseen | – | |
| write-through-pointer | 11 | unseen | – | |
| nil-pointer-guard | 11 | unseen | – | |
| pointer-receivers (mutate, consistency) | 11 | unseen | – | |
| method-sets-and-interfaces (*T vs T) | 11 | unseen | – | |
| pointer-passed-by-value (reassign vs write) | 11 | unseen | – | |
| pointers-in-maps-structs | 11 | unseen | – | |
| package-main-vs-library | 12 | unseen | – | |
| exported-capitalisation | 12 | unseen | – | |
| package-per-directory | 12 | unseen | – | |
| go-run-build-install | 12 | unseen | – | |
| stale-binary | 12 | unseen | – | |
| go-mod-replace-work | 12 | unseen | – | |
| import-paths (module + subdir) | 12 | unseen | – | |
| go-get-tidy | 12 | unseen | – | |
| clean-package-api | 12 | unseen | – | |
| goroutine-go-keyword | 13 | unseen | – | |
| channel-send-receive (blocking) | 13 | unseen | – | |
| deadlock-diagnosis | 13 | unseen | – | |
| token-channel (struct{}) | 13 | unseen | – | |
| buffered-channels | 13 | unseen | – | |
| close-and-ok | 13 | unseen | – | |
| range-over-channel | 13 | unseen | – | |
| select (ok checks) | 13 | unseen | – | |
| select-default-nonblocking | 13 | unseen | – | |
| channel-direction-types | 13 | unseen | – | |
| mutex-lock-defer-unlock | 14 | unseen | – | |
| race-condition-recognition | 14 | unseen | – | |
| rwmutex-readers-writers | 14 | unseen | – | |
| mutex-copy-and-reentrancy | 14 | unseen | – | |
| race-detector | 14 | unseen | – | |
| type-parameters-syntax | 15 | unseen | – | |
| zero-value-of-T | 15 | unseen | – | |
| constraints-any-vs-methods | 15 | unseen | – | |
| type-lists-ordered-tilde | 15 | unseen | – | |
| parametric-interfaces | 15 | unseen | – | |
| generic-types-methods | 15 | unseen | – | |
| type-inference-explicit | 15 | unseen | – | |
| type-definitions (distinct types, conversion) | 16 | unseen | – | |
| iota-basics (repeat rule, reset per block) | 16 | unseen | – | |
| iota-skip-offset (_ first, shifts) | 16 | unseen | – | |
| enum-string-method | 16 | unseen | – | |
| enum-validation | 16 | unseen | – | |

## Struggle Patterns
- (durable patterns only — "python mind" specifics go here once observed)
- **int division / conversion timing** (seen 09-12, 09-13, 09-14): expects `/` to keep the fraction; `float64(a / b)` looked like it should rescue it. Keep seeding PREDICTs where the conversion comes too late. 09-14: explained the truncation-before-conversion bug correctly without help, so the model looks to be landing.
- Watch: fixes a type mismatch by changing the function signature (int params → float64) instead of converting inside the function, seen 2026-09-14. Understood why after talking about constants (no type until used) versus variables (type fixed when created).
- Watch: python type-hint reflex on declarations (`x T := v`), seen 2026-09-12, not yet durable.
- Watch: expects `Println` of a whole float64 to show `11.0` (Python repr), seen 2026-09-13.

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

## Session Log
- 2026-09-09: dojo created; rep 001 loaded.
- 2026-09-12 (session 1): reps 001–003 at 2/10, all pass-with-notes; rep 004 loaded, paused for the night.
- 2026-09-13 (session 2): rep 004 clean; 005 pass-notes; rep 006 loaded, paused before attempt.
- 2026-09-14 (session 3): rep 006 pass-notes at 3/10 after one nudge; discussed parameter order versus positional return values and why `members, pot` works in the zero-members guard. Rep 007 loaded, paused before attempt; scope stays ch1–4 at 2–3.
- 2026-09-14 (session 4): rep 007 clean (blind PREDICT, shadowing); 008 pass-notes (fixed int division by changing the signature). Rep 009 loaded, paused before attempt. Talked about untyped constants versus typed variables (why float64 params only worked with literal arguments). Scope stays ch1–4 at 2–3.
- 2026-09-12: original rep 001 (splitName) withdrawn before attempt — relied on slicing (ch9) and strings.Index Matt hasn't met. Replaced with a ch1 FIX at 2/10.
