package main

import "fmt"

// tally hands back a running-total function for one counter.
func tally() func(int) int {
	sold := 0
	return func(n int) int {
		sold += n
		return sold
	}
}

func main() {
	rolls := tally()
	fmt.Println(rolls(12))
	fmt.Println(rolls(8))
	fmt.Println(rolls(30))

	pies := tally()
	fmt.Println(pies(5))
	fmt.Println(rolls(1))
}
