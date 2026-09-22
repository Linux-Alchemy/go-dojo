package main

import "fmt"

type parcel struct {
	weightKg int
	express  bool
}

func (p parcel) cost() int {
	price := 4
	if p.weightKg > 5 {
		price += (p.weightKg - 5) * 2
	}
	if p.express {
		price *= 2
	}
	return price
}

func main() {
	fmt.Println(parcel{weightKg: 3}.cost())
	fmt.Println(parcel{weightKg: 5}.cost())
	fmt.Println(parcel{weightKg: 8}.cost())
	fmt.Println(parcel{weightKg: 8, express: true}.cost())
	fmt.Println(parcel{express: true}.cost())
}
