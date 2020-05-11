package main

import "fmt"

func main() {
	beatles := map[string]string{
		"John":   "guitar",
		"Paul":   "bass",
		"George": "guitar",
		"Ringo":  "drums",
	}

	delete(beatles, "John")

	fmt.Println(beatles)
	// map[Paul:bass George:guitar Ringo:drums]
}
