package main

import "fmt"

func main() {
	tickets := 8
	if tickets > 0 {
		tickets := tickets - 3
		fmt.Println("inside:", tickets)
	}
	fmt.Println("after:", tickets)
	tickets = tickets - 2
	fmt.Println("final:", tickets)
}
