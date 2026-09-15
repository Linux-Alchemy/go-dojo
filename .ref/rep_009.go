package main

import "fmt"

func chargeFor(minutes int) (int, string) {
	switch {
	case minutes <= 30:
		return 0, "grace"
	case minutes <= 120:
		return 4, "short"
	case minutes <= 480:
		return 12, "day"
	default:
		return 20, "overnight"
	}
}

func report(minutes int) {
	fee, band := chargeFor(minutes)
	fmt.Printf("%d min: %s, $%d\n", minutes, band, fee)
}

func main() {
	report(30)
	report(31)
	report(120)
	report(480)
	report(481)
}
