package main

import (
	"fmt"
	"strings"
)

// section: main
// slice[starting_index : (starting_index + length)]

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}

	fmt.Println(names) // [John Paul George Ringo]

	// Get the first three elements of the `names` slice
	guitars := names[:3]

	fmt.Println(guitars) // [John Paul George]

	for i, g := range guitars {
		guitars[i] = strings.ToUpper(g)
	}

	// Print out our original slice of names
	fmt.Println(names) // [JOHN PAUL GEORGE Ringo]
}

// section: main
