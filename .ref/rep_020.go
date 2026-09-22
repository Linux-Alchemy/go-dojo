package main

import "fmt"

type stats struct {
	hp    int
	level int
}

type hero struct {
	name  string
	stats stats
}

func main() {
	a := hero{name: "Wren", stats: stats{hp: 20, level: 3}}
	b := a
	b.name = "Moss"
	b.stats.hp = 5

	fmt.Println(a.name, a.stats.hp)
	fmt.Println(b.name, b.stats.hp, b.stats.level)

	var c hero
	c.stats.level++
	fmt.Printf("%q %d %d\n", c.name, c.stats.hp, c.stats.level)

	a.stats = b.stats
	b.stats.level = 9
	fmt.Println(a.stats.hp, a.stats.level)
}
