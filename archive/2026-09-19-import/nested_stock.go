//go:build ignore

package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 014 · ch5 structs · difficulty 2/10 · ~2 min         ║
// ║ shape: FIX · tags: nested-structs                            ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  A shop keeps stock details inside each record.     ║
// ║ This validator does not compile.                            ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                       ║
// ║ 1. Fix validStock without changing the structs or calls.    ║
// ║ 2. A record needs a non-empty name and a count of zero or   ║
// ║    more. The featured flag can be either true or false.     ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                            ║
// ║ in stock: true                                             ║
// ║ sold out: true                                             ║
// ║ negative count: false                                      ║
// ║ missing name: false                                        ║
// ║ empty record: false                                        ║
// ╚════════════════════════════════════════════════════════════════╝

type stockDetails struct {
	name   string
	onHand int
}

type stockRecord struct {
	details  stockDetails
	featured bool
}

func validStock(item stockRecord) bool {
	return item.details.name != "" && item.details.onHand >= 0
}

func main() {
	fmt.Println("in stock:", validStock(stockRecord{details: stockDetails{name: "Bolts", onHand: 12}, featured: true}))
	fmt.Println("sold out:", validStock(stockRecord{details: stockDetails{name: "Nuts"}}))
	fmt.Println("negative count:", validStock(stockRecord{details: stockDetails{name: "Washers", onHand: -2}}))
	fmt.Println("missing name:", validStock(stockRecord{details: stockDetails{onHand: 5}}))
	fmt.Println("empty record:", validStock(stockRecord{}))
}

/*
Verdict: clean. Corrected nested field access on first check without nudges.
All five outputs match; gofmt and go vet clean. No idiom flags.
*/
