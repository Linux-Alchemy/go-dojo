package main

import "fmt"

var spaces = 120

func main() {
	taken := 87
	var rate float64 = 2.50
	free := spaces - taken
	fmt.Printf("Free spaces: %d of %d\n", free, spaces)
	fmt.Printf("Hourly rate: £%.2f\n", rate)
}
