package main

import "fmt"

// ticketPrice returns the price in dollars for one seat at the Roxy.
func ticketPrice(age int) int {
	if age < 16 {
		return 7
	}
	if age < 65 {
		return 12
	}
	return 9
}

func main() {
	fmt.Printf("age 8: $%d\n", ticketPrice(8))
	fmt.Printf("age 15: $%d\n", ticketPrice(15))
	fmt.Printf("age 16: $%d\n", ticketPrice(16))
	fmt.Printf("age 40: $%d\n", ticketPrice(40))
	fmt.Printf("age 64: $%d\n", ticketPrice(64))
	fmt.Printf("age 65: $%d\n", ticketPrice(65))
	fmt.Printf("age 80: $%d\n", ticketPrice(80))
}
