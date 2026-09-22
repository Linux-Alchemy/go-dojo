package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 021 · ch5 structs · difficulty 3/10  ·  ~4 min        ║
// ║ shape: COMPLETE · tags: methods-value-receiver                 ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  Dash & Dump Couriers price every parcel the same way, ║
// ║ and the till wants to ask the parcel itself what it costs.     ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ TASK                                                           ║
// ║ 1. Write a method called cost on parcel. It returns the price  ║
// ║    in whole dollars:                                           ║
// ║      · $4 flat for any parcel                                  ║
// ║      · plus $2 for every kg over 5 kg                          ║
// ║      · express doubles the whole price                         ║
// ║ 2. Once it exists, uncomment the prints in main.               ║
// ║ 3. Don't touch the struct or the parcels.                      ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ EXPECTED OUTPUT                                                ║
// ║   4                                                            ║
// ║   4                                                            ║
// ║   10                                                           ║
// ║   20                                                           ║
// ║   8                                                            ║
// ╚════════════════════════════════════════════════════════════════╝

type parcel struct {
	weightKg int
	express  bool
}

func main() {
	letter := parcel{weightKg: 3}
	boots := parcel{weightKg: 5}
	kettle := parcel{weightKg: 8}
	urgentKettle := parcel{weightKg: 8, express: true}
	urgentEnvelope := parcel{express: true}

	_, _, _, _, _ = letter, boots, kettle, urgentKettle, urgentEnvelope // keeps the compiler quiet until the prints are live; delete it then

	// fmt.Println(letter.cost())
	// fmt.Println(boots.cost())
	// fmt.Println(kettle.cost())
	// fmt.Println(urgentKettle.cost())
	// fmt.Println(urgentEnvelope.cost())
	fmt.Print() // keeps the fmt import used until the prints are live; delete it then
}
