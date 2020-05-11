package main

import (
	"fmt"
)

// slice[starting_index : (starting_index + length)]

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}

	fmt.Println(names) // [John Paul George Ringo]

	// Get 2 elements starting with the second element (index 1)
	fmt.Println(names[1:3]) // [Paul George] - names[1:1+2]

	// functionally equivalent
	fmt.Println(names[2:len(names)]) // [George Ringo]
	fmt.Println(names[2:])           // [George Ringo]

	// functionally equivalent
	fmt.Println(names[0:2]) // [John Paul]
	fmt.Println(names[:2])  // [John Paul]
}
