# Go Dojo

Lives at `~/github/go-dojo`. Open Claude Code here and say `obie-wan` + `dojo`. The reference notes stay in the grimoire under `learn/Go/`; nothing there is edited during practice.

Shared scratch pad for short Go reps. Matt opens `main.go` in nvim; Claude (via the `go-dojo` skill) writes a challenge card into it; Matt writes the solution underneath; Claude checks it, logs it, and writes the next one.

## Layout

| Path | What |
|---|---|
| `main.go` | **The pad.** One rep at a time. Card at the top, your code below. |
| `go.mod` | So `go run .` works. Module name `dojo`. Leave it alone. |
| `DOJO_LOG.md` | The ledger: concept ratings, rep history. Claude maintains it. |
| `SESSIONS.md` | Append-only session notes: scope, average difficulty, solid, struggled, coach's read. Grows forever; that's the point. |
| `archive/` | Every completed rep, as `rep_NNN_<slug>.go` with `//go:build ignore` so it never compiles by accident. Your solution, verbatim, plus Claude's verdict as a trailing comment. |
| `.ref/` | Claude's reference solutions, used to verify expected output before a rep is published. Not for peeking. |

## Commands you'll use

```bash
go run .
```

```bash
gofmt -l . && go vet ./...
```

## Vocabulary

| Say | Happens |
|---|---|
| `dojo, ch3 and ch4, 3–5` | You name topics and a difficulty range. Reps are random inside that. Claude writes the first card into `main.go`. |
| `next rep` | Another one, same scope. |
| `check` / `done` | Claude runs gofmt, vet, `go run`, diffs output against the card, reviews idioms, gives a verdict, archives, logs, writes the next rep. |
| `stuck` / `nudge` | One-sentence pointer. No code. Climbs the ladder only if asked again. |
| `show me` | Reference solution pasted into chat (not the file), rep logged as **revisit**. |
| `easier` / `harder` / `ch3 only` / `mix` | Steers selection for the rest of the session. |
| `that's enough for today` / `close` | Ledger updated, session note appended to `SESSIONS.md`, three-line summary, and a stagnation call-out if one is due. |

## Rep types

`FIX` complete program with a planted bug · `COMPLETE` starter code plus numbered instructions (the Boot.dev default) · `PREDICT` write the output before running · `TRANSLATE` Python → Go. Imports and anything you haven't been taught yet are always pre-supplied.

Difficulty is the Boot.dev 1–10 scale. Time box is 2–8 minutes. If a rep blows past 8 minutes, say so and it gets logged as **revisit** rather than ground out.
