package main

import "fmt"

// ╔════════════════════════════════════════════════════════════════╗
// ║ DOJO REP 023 · ch5 structs · difficulty 3/10 · ~4 min         ║
// ║ shape: TRANSLATE · tags: anonymous-structs                    ║
// ╠════════════════════════════════════════════════════════════════╣
// ║ CONTEXT  A pub lists how many seats a booth can offer.       ║
// ║ TASK                                                        ║
// ║ 1. Translate the Python below into Go inside main.          ║
// ║ 2. Use an anonymous struct literal for booth, with fields   ║
// ║    matching the dictionary keys. Choose their Go types.     ║
// ║    Do not declare a named struct type or use a map.         ║
// ║ 3. Keep the data and behaviour; show is supplied for you.   ║
// ║ EXPECTED OUTPUT                                             ║
// ║   Moon 0                                                    ║
// ╚════════════════════════════════════════════════════════════════╝

// Python:
// booth = {"name": "Moon", "seats": 8, "reserved": True}
// available = booth["seats"]
// if booth["reserved"]:
//     available = 0
// show(booth["name"], available)

// show prints the booth name and the number of seats available.
func show(name string, available int) {
	fmt.Println(name, available)
}

func main() {
	// Your translation here.
}
