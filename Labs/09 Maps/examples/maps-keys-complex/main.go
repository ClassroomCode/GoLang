package main

import (
	"fmt"
)

// section: key
type simple struct {
	ID int
}

type complex struct {
	f func(id int) simple
}

func main() {
	m := map[simple]string{}

	fmt.Println(m)

	// invalid map key type complex
	m1 := map[complex]string{}
}

// section: key
