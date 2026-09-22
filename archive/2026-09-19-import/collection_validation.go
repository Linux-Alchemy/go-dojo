//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 013 · ch5 structs · difficulty 3/10 · ~4 min         ║
// ║ shape: COMPLETE · tags: typed-value-validation,              ║
// ║ struct-zero-values, nested-structs                           ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  A shop checks whether an order can be collected.    ║
// ║ emailReceipt records the customer's receipt preference.     ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                       ║
// ║ 1. Complete canCollect; keep the types and calls supplied.   ║
// ║ 2. Collection requires a non-empty customer name and a      ║
// ║    non-zero pickup code. Those are the only requirements.   ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                            ║
// ║ ready: true                                                ║
// ║ paper receipt: true                                        ║
// ║ missing name: false                                        ║
// ║ missing code: false                                        ║
// ║ empty order: false                                         ║
// ╚════════════════════════════════════════════════════════════════╝

type contact struct {
	name       string
	pickupCode int
}

type collection struct {
	customer     contact
	emailReceipt bool
}

func canCollect(order collection) bool {
	if order.customer.name != "" && order.customer.pickupCode != 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println("ready:", canCollect(collection{customer: contact{name: "Ada", pickupCode: 42}, emailReceipt: true}))
	fmt.Println("paper receipt:", canCollect(collection{customer: contact{name: "Ben", pickupCode: 17}}))
	fmt.Println("missing name:", canCollect(collection{customer: contact{pickupCode: 42}}))
	fmt.Println("missing code:", canCollect(collection{customer: contact{name: "Cal"}}))
	fmt.Println("empty order:", canCollect(collection{}))
}

/*
Verdict: pass, notes. All five outputs match; gofmt and go vet clean.
Flag: the if returning true followed by return false can be a direct return of the boolean expression.
No nudges; value rules correctly checked and optional false preference accepted.
*/
