# Ch4 · Currying

Boot.dev lesson, added 2026-09-17. Last lesson of chapter 4. Matt completed it but found it quite a bit more challenging than closures. Ledger tags: `currying (func returning func)`, leaning on `higher-order (func params, func types)`.

## Lesson text

Function currying is a concept from functional programming and involves partial application of functions. It allows a function with multiple arguments to be transformed into a sequence of functions, each taking a single argument.

Let's simulate this behavior. For example:

```go
func main() {
	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(add)

	fmt.Println(squareFunc(5))
	// prints 25

	fmt.Println(doubleFunc(5))
	// prints 10
}

func multiply(x, y int) int {
	return x * y
}

func add(x, y int) int {
	return x + y
}

func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}
```

In the example above:

- `selfMath(multiply)` returns a new function and stores it in `squareFunc`; it doesn't run `multiply` yet.
- `squareFunc(5)` runs the returned function, which calls `multiply(5, 5)`.

## Assignment

The Textio API needs a very robust error-logging system so we can see when things are going awry in the back-end system. We need a function that can create a custom "logger" (a function that prints to the console) given a specific formatter.

These errors are test data, not runtime failures.

Complete the `getLogger` function. It should take as input a formatter function and return a new function. The new logger function takes as input two strings and passes them to the formatter, then prints the result. Keep the order of the strings.

## Matt's solution (passed)

The challenge was to fill in the body of `getLogger`. Everything else was supplied.

```go
package main

import (
	"errors"
	"fmt"
)

// getLogger takes a function that formats two strings into
// a single string and returns a function that formats two strings but prints
// the result instead of returning it
func getLogger(formatter func(string, string) string) func(string, string) {
	return func(first, second string) {
		fmt.Println(formatter(first, second))
	}
}

// don't touch below this line

func test(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("====================================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func colonDelimit(first, second string) string {
	return first + ": " + second
}

func commaDelimit(first, second string) string {
	return first + ", " + second
}

func main() {
	dbErrors := []error{
		errors.New("out of memory"),
		errors.New("cpu is pegged"),
		errors.New("networking issue"),
		errors.New("invalid syntax"),
	}
	test("Error on database server", dbErrors, colonDelimit)

	mailErrors := []error{
		errors.New("email too large"),
		errors.New("non alphanumeric symbols found"),
	}
	test("Error on mail server", mailErrors, commaDelimit)
}
```

## What to drill

- **Splitting the signature.** `func getLogger(formatter func(string, string) string) func(string, string)` has three parts: the name, one parameter called `formatter` whose type is `func(string, string) string`, and a return type of `func(string, string)`. Finding where the parameter list closes is the whole trick.
- **The returned function returns nothing.** `func(string, string)` has no result type, so the inner function prints and has no `return` line. The formatter's type differs from the logger's by exactly that trailing `string`.
- **Passing a function versus calling it.** `getLogger(colonDelimit)` has no brackets after `colonDelimit`, so it hands the function over without running it. `formatter(first, second)` has brackets, so it runs.
- **It's a closure underneath.** The inner function keeps hold of `formatter` after `getLogger` has returned, the same way `adder` kept `totalTexts` (see `ch4_closures.md`).
- **Setup once, use many times.** `getLogger` runs once per `test` call. `logger` runs once per error.
- **The inner function's signature is dictated.** It must match the declared return type exactly: two strings in, nothing out. Parameter names are free; types and count are not.

## Reps: what to pre-supply

The Boot.dev scaffold uses `errors.New`, `[]error`, `for ... range` and `defer`, none of which Matt has drilled (ch7, ch8, ch9, and the ch4 defer tag is unseen). Reps on this tag should keep `main` to plain calls and leave that machinery out, or supply it whole and out of the way.

## Rep seeds

Build up in this order; the first two are `higher-order` on its own, before anything gets returned.

- PREDICT: a function that takes `func(int, int) int` and calls it with fixed arguments; pass it `add`, then `multiply`.
- COMPLETE: write `apply(a, b int, op func(int, int) int) int`, signature supplied.
- PREDICT: `selfMath`-style factory, trace which function runs when, including a line that builds a function and never calls it.
- FIX: the factory calls the formatter straight away and returns the wrong thing, so it won't compile against the declared return type.
- FIX: the inner function swaps the two strings (the "keep the order" clause).
- COMPLETE: `getLogger` cousin with the signature supplied, e.g. a `getPricer(taxFunc func(float64) float64) func(float64)` that prints the taxed price.
- COMPLETE, harder: same again, but Matt writes the signature from a plain-English description.
