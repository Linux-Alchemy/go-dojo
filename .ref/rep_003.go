package main

import "fmt"

const hoursOpen = 11
const minutesPerHour = 60
const songMinutes = 3
const songsPerDay = hoursOpen * minutesPerHour / songMinutes
const pricePerSong = 0.50

func main() {
	fmt.Printf("Songs per day: %d\n", songsPerDay)
	fmt.Printf("Takings: £%.2f\n", songsPerDay*pricePerSong)
}
