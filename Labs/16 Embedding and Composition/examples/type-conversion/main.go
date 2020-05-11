package main

import (
	"fmt"
)

// section: code
func main() {
	var i uint64 = 10
	var j int = 20

	// We must type convert i from a uint64 to an int in order to add i and j
	k := (int(i) + j)

	fmt.Println(k)
}

// section: code
