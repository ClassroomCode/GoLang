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
	key := "Paul"

	if value, ok := beatles[key]; ok {
		fmt.Printf("Found key %q: %q", key, value) // Found Key "Paul": "bass"
	} else {
		fmt.Printf("Key not found: %q", key) // Key not found: "foo"
	}

	// Change the "key" to "foo" and re-run the example
	// section: code
}
