package main

import "fmt"

type ticket struct {
	plate string
	bay   string
	hours int
	rate  int
}

func describe(t ticket) string {
	return fmt.Sprintf("%s in bay %s: %dh at $%d/h = $%d", t.plate, t.bay, t.hours, t.rate, t.hours*t.rate)
}

func main() {
	a := ticket{plate: "KX19 ABC", bay: "B12", hours: 3, rate: 2}
	b := ticket{plate: "LM70 XYZ", bay: "C04", hours: 2, rate: 5}
	c := ticket{plate: "PQ21 RST", bay: "A01", hours: 1, rate: 4}

	fmt.Println(describe(a))
	fmt.Println(describe(b))
	fmt.Println(describe(c))
}
