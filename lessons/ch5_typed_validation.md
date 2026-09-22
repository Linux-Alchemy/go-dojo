# Ch5 · Nested structs and typed validation

Boot.dev assignment imported 2026-09-19. Matt supplied the solution below and identified the difficulty as the mental model, not syntax. No dojo rep or independent course-pass result is claimed by this import.

## Assignment (as supplied)

Textio has a bug, we've been sending texts that are missing critical bits of information! Before we send text messages in Textio, we must check to make sure the required fields have non-zero values.

Notice that the user struct is a nested struct within the messageToSend struct. Both sender and recipient are user struct types.

Complete the canSendMessage function. It should return true only if the sender and recipient fields each contain a name and a number. If any of the default zero values are present, return false instead.

## Matt's supplied solution

```go
package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
		return false
	}
	return true
}
```

## Learner evidence

Matt originally tried to check all possible kinds of value a field could hold, including whether a declared string might be an int or bool. He identified that Go's static types already rule those out at compilation. He reports that differences in the Go/Python mental model have obstructed several otherwise straightforward lessons.

The supplied solution satisfies the stated conditions. The message body is intentionally not validated: the assignment only requires sender and recipient names and numbers. Negative numbers are non-zero; rejecting them would add an unstated rule.

## What to drill

- Separate guarantees supplied by the declared concrete type from value rules supplied by the problem.
- A concrete string field holds a string; test its permitted values, not whether it has turned into a bool or int. Python values also have types, but ordinary annotations do not enforce assignment types at runtime.
- An omitted struct field still exists with its type's zero value; nested structs recursively contain zero-valued fields.
- Zero is not universally invalid: false may be a legitimate preference, and zero may be a valid count. Read the contract before rejecting it.
- Validate only required fields; do not invent extra constraints.
- Keep mental-model differences visible across reps, using observed mistakes rather than attributing every error to Python.

## Rep seeds

- COMPLETE: nested collection details require a contact name and pickup code; an optional boolean preference may be false. Structs and calls supplied.
- FIX: a validator incorrectly rejects an optional false flag or a permitted zero count.
- PREDICT: omitted fields versus explicit zero values in nested struct literals; trace the validator's result.

Tags: typed-value-validation (revisit from learner report), struct-zero-values and nested-structs (unseen until drilled). No mastery inferred from supplied code.
