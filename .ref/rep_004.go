package main

import "fmt"

// postage returns the parcel size band for a weight in kg.
func postage(weightKg float64) string {
	if weightKg <= 0 {
		return "invalid"
	}
	if weightKg <= 2 {
		return "small"
	}
	if weightKg <= 10 {
		return "medium"
	}
	return "large"
}

func main() {
	fmt.Printf("%.2fkg: %s\n", 0.0, postage(0))
	fmt.Printf("%.2fkg: %s\n", 1.5, postage(1.5))
	fmt.Printf("%.2fkg: %s\n", 2.0, postage(2))
	fmt.Printf("%.2fkg: %s\n", 9.99, postage(9.99))
	fmt.Printf("%.2fkg: %s\n", 10.0, postage(10))
	fmt.Printf("%.2fkg: %s\n", 10.5, postage(10.5))
}
