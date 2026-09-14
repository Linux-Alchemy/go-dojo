package main

import "fmt"

func main() {
	tab := 45
	mates := 4

	share := tab / mates
	exact := float64(tab) / float64(mates)
	oops := float64(tab / mates)

	fmt.Println(share)
	fmt.Println(share * mates)
	fmt.Printf("%.2f\n", exact)
	fmt.Printf("%.2f\n", oops)
	fmt.Println(oops)
	fmt.Println(exact)
}
