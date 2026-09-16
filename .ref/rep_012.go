package main

import "fmt"

func pickMessage(orderID string, qty int, inStock int) string {
	if qty <= 0 {
		return fmt.Sprintf("%s: nothing to pick", orderID)
	}
	if inStock < qty {
		short := qty - inStock
		noun := "items"
		if short == 1 {
			noun = "item"
		}
		return fmt.Sprintf("%s: short by %d %s", orderID, short, noun)
	}
	noun := "items"
	if qty == 1 {
		noun = "item"
	}
	return fmt.Sprintf("%s: pick %d %s", orderID, qty, noun)
}

func main() {
	fmt.Println(pickMessage("A101", 3, 10))
	fmt.Println(pickMessage("A102", 1, 5))
	fmt.Println(pickMessage("A103", 4, 3))
	fmt.Println(pickMessage("A104", 6, 6))
	fmt.Println(pickMessage("A105", 0, 9))
	fmt.Println(pickMessage("A106", 5, 2))
}
