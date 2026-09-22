package main

import "fmt"

type keg struct {
	beer    string
	litres  int
	chilled bool
}

func main() {
	bitter := keg{beer: "Bitter", litres: 50, chilled: true}
	stout := keg{beer: "Stout", litres: 30}
	mild := keg{beer: "Mild"}
	empty := keg{}

	fmt.Printf("%+v\n", bitter)
	fmt.Printf("%+v\n", stout)
	fmt.Printf("%+v\n", mild)
	fmt.Printf("%+v\n", empty)
}
