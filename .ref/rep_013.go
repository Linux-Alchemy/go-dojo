package main

import "fmt"

// shelfNote returns the librarian's note for one book.
func shelfNote(title string, daysOut, loanDays int) string {
	late := daysOut - loanDays
	fine := 1
	if late > 7 {
		fine = 3
	}
	if daysOut == 0 {
		return title + ": on the shelf"
	}
	if late > 0 {
		return fmt.Sprintf("%s: %d days late, fine $%d", title, late, fine)
	}
	return fmt.Sprintf("%s: %d days left", title, -late)
}

func main() {
	fmt.Println(shelfNote("Dune", 3, 14))
	fmt.Println(shelfNote("Emma", 20, 14))
	fmt.Println(shelfNote("Ulysses", 30, 14))
	fmt.Println(shelfNote("Heidi", 0, 14))
	fmt.Println(shelfNote("Kim", 14, 14))
	fmt.Println(shelfNote("Momo", 21, 14))
}
