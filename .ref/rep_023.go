package main

import "fmt"

// show prints the booth name and the number of seats available.
func show(name string, available int) {
	fmt.Println(name, available)
}

func main() {
	booth := struct {
		name     string
		seats    int
		reserved bool
	}{
		name:     "Moon",
		seats:    8,
		reserved: true,
	}
	available := booth.seats
	if booth.reserved {
		available = 0
	}
	show(booth.name, available)
}
