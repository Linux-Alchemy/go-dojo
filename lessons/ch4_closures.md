# Ch4 · Closures

Boot.dev lesson, added 2026-09-17. Matt completed it but flagged it as one that jammed him up and needs review. Ledger tag: `closures (capture by reference)`.

## Lesson text

A closure is a function that references variables from outside its own function body. The function may access and assign to the referenced variables.

In this example, the `concatter()` function returns a function that has reference to an enclosed `doc` value. Each successive call to `harryPotterAggregator` mutates that same `doc` variable.

```go
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func main() {
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive
}
```

## Assignment

Keeping track of how many texts we send is mission-critical at Textio. Complete the `adder()` enclosing function.

1. Create an enclosed `sum` value inside the `adder()` function.
2. Return a function from the `adder()` function that adds its input (an `int`) to the `sum` and returns the new value of `sum`. (In other words, it keeps a running total of the `sum` variable within a closure.)

## Matt's solution (passed)

```go
package main

func adder() func(int) int {
	totalTexts := 0
	return func(textMsg int) int {
		totalTexts += textMsg
		return totalTexts
	}
}
```

## What to drill

- **Reading the signature.** `func adder() func(int) int` is "adder takes nothing and returns a `func(int) int`". The second `func(int) int` is a type, the same way `string` would be.
- **Two calls, two jobs.** `adder()` runs once and sets up `totalTexts`. The function it hands back runs every time after that. `totalTexts := 0` does not run again.
- **The variable outlives the call.** `totalTexts` is local to `adder`, but the returned function still holds it, so it survives after `adder` returns. The inner function uses the variable itself, not a copy of its value at the time.
- **One closure, one variable.** Each call to `adder()` makes a fresh `totalTexts`. Two adders keep two separate totals.
- **`+=` versus `:=` inside.** `totalTexts += n` updates the enclosed variable. `totalTexts := ...` inside the inner function would make a new one and the total would never move (shadowing, rep 007).

## Rep seeds

- PREDICT: two counters from one factory, calls interleaved, print each result.
- FIX: inner function shadows the enclosed variable with `:=`, so the total stays put.
- FIX: the enclosed variable is declared inside the inner function, so it resets every call.
- COMPLETE: a running-maximum or running-average closure, signature supplied.
