package main

import "fmt"

// scorer hands back a running-total function for one darts player.
func scorer() func(int) int {
	total := 0
	return func(points int) int {
		total += points
		return total
	}
}

func main() {
	alex := scorer()
	sam := scorer()

	fmt.Println(alex(20))
	fmt.Println(sam(5))
	fmt.Println(alex(19))
	fmt.Println(sam(1))
	fmt.Println(alex(60))

	bailey := scorer()
	fmt.Println(bailey(7))
	fmt.Println(sam(3))
}
