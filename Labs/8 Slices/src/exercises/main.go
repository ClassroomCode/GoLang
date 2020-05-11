package main

import "fmt"

func main() {
	parents := []string{"Carol", "Mike"}
	kids := []string{"Marcia", "Jan", "Cindy", "Greg", "Peter", "Bobby"}

	// TODO: Create a new slice called family by joining the parents and kids slice together

	// TODO: Print out the value of `family`

	// TODO: Print the length and capacity of the new slice.

	// TODO: Create a new slice called "youngest" by taking the last 3 elements of the slice you just made.
	// HINT: slice[starting_index : (starting_index + length)], or use the shorthand version of slicing

	// TODO: Print out the value of the new slice `youngest` that you just created

	// TODO: Fix the following bugs
	extras := make([]string, 1, 0)
	extras[1] = "Alice"
	fmt.Println(extras)
}
