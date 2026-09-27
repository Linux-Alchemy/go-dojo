package main

import "fmt"

type team struct {
	name    string
	captain string
}

type entry struct {
	team
	round int
	score int
}

func main() {
	e := entry{
		team:  team{name: "Quizzly Bears", captain: "Ada"},
		round: 3,
		score: 7,
	}
	fmt.Println(e.name, "captained by", e.captain)

	total := e.score
	if e.score >= 7 {
		total = total + 3
	}
	fmt.Printf("round %d: %d points\n", e.round, total)
}
