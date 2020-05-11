package main

import "fmt"

func main() {
	beatles := map[string]string{
		"John":   "guitar",
		"Paul":   "bass",
		"George": "guitar",
		"Ringo":  "drums",
	}

	// section: code
	for key, value := range beatles {
		fmt.Printf("%s plays %s\n", key, value)
	}

	// John plays guitar
	// Paul plays bass
	// George plays guitar
	// Ringo plays drums
	// section: code
}
