package main

import "fmt"

func main() {
	pizzas := 3
	teams := 7
	slices := pizzas * 8
	each := float64(slices) / float64(teams)

	fmt.Println(slices / teams)
	fmt.Printf("%.1f\n", each)
	fmt.Printf("%T %T\n", slices, each)
	fmt.Println(slices, "slices for", teams, "teams")
}
