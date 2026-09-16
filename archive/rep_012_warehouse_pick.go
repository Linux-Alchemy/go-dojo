//go:build ignore

package main

import (
	"fmt"

)

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 012 · ch2/ch3/ch4 · difficulty 3/10 · ~5 min          ║
// ║ shape: COMPLETE · tags: no-ternary, guard-clauses, sprintf     ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  The warehouse handheld shows pickers one line per     ║
// ║ order. The old Python used `"item" if n == 1 else "items"`.    ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK  Write the body of pickMessage so it returns (not         ║
// ║ prints) the line for an order:                                 ║
// ║ 1. An order for zero or fewer items has nothing to pick.       ║
// ║ 2. If the shelf can't cover the order, say how many short.     ║
// ║ 3. Otherwise, say how many to pick.                            ║
// ║ 4. Counts of one read "item", every other count "items".       ║
// ║ 5. Don't touch main.                                           ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║  A101: pick 3 items                                            ║
// ║  A102: pick 1 item                                             ║
// ║  A103: short by 1 item                                         ║
// ║  A104: pick 6 items                                            ║
// ║  A105: nothing to pick                                         ║
// ║  A106: short by 3 items                                        ║
// ╚════════════════════════════════════════════════════════════════╝

// pickMessage returns the handheld's line for one order.
func pickMessage(orderID string, qty int, inStock int) string {
	if qty > inStock {
		missingStock := qty - inStock
		if missingStock == 1 {
			return fmt.Sprintf("%s: short by %d item", orderID, missingStock)
		} else {
			return fmt.Sprintf("%s: short by %d items", orderID, missingStock)
		}
	}
	if qty <= 0 {
		return fmt.Sprintf("%s: nothing to pick", orderID) 
	}
	if qty == 1 {
		return fmt.Sprintf("%s: pick %d item", orderID, qty)
	}else {
		return fmt.Sprintf("%s: pick %d items", orderID, qty)
	}

}

func main() {
	fmt.Println(pickMessage("A101", 3, 10))
	fmt.Println(pickMessage("A102", 1, 5))
	fmt.Println(pickMessage("A103", 4, 3))
	fmt.Println(pickMessage("A104", 6, 6))
	fmt.Println(pickMessage("A105", 0, 9))
	fmt.Println(pickMessage("A106", 5, 2))
}

/*
VERDICT: pass, notes (several nudges → tags to revisit)

- All six lines matched and vet was clean, after four rounds of help: stacked ifs overwriting one variable,
  an if/else that returned on both branches (unreachable code below it), treating
  `missingStock := qty - inStock` as though it were a check, and finally where the shortfall check goes.
- Final shape is right: shortfall check on the outside, item/items choice inside it, each rule returns.

FLAGS
1. gofmt dirty: blank line inside the import block, trailing space after the nothing-to-pick Sprintf,
   `}else`, blank line before the closing brace. (gopls also auto-added golang.org/x/text/message earlier.)
2. `else` after a branch that returns isn't needed. Each item/items pair also repeats the whole Sprintf.
   Default-then-override (`noun := "items"; if n == 1 { noun = "item" }`) needs only one.
3. The `qty <= 0` guard runs second. It works here, but degenerate cases usually go first so nobody
   has to prove a zero order can never count as short.
*/
