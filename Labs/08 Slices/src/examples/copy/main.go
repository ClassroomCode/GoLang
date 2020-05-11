package main

import "fmt"

func main() {
	original := []string{"carrot", "potato", "cucumber", "onion"}
	// create a new reference to the existing slice
	ref := original

	// initialize a variable named 'dup' to the
	// same size as the original slice
	dup := make([]string, len(original))

	// copy the values from original to dup
	// NOTE: if the slices are not the same size, it will
	// only copy what it has space for (length)
	copy(dup, original)

	// changing either `ref` or `original` will change either
	// of them as they still share the reference to the same
	// backing array
	ref[0] = "zuchinni"
	original[1] = "tomato"

	fmt.Println("original: ", original)
	fmt.Println("ref:      ", ref)
	fmt.Println("dup:      ", dup)
}
