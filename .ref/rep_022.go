package main

import "fmt"

type lantern struct {
	fuel int
}

func (l lantern) minutes() int {
	if l.fuel <= 0 {
		return 0
	}
	return l.fuel * 3
}

func main() {
	empty := lantern{}
	full := lantern{fuel: 6}
	fmt.Println(empty.minutes())
	fmt.Println(full.minutes())
}
