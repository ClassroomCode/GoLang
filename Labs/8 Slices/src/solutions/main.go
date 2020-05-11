package main

import "fmt"

func main() {
	parents := []string{"Carol", "Mike"}
	kids := []string{"Marcia", "Jan", "Cindy", "Greg", "Peter", "Bobby"}

	// TODO: Create a new slice called family by joining the parents and kids slice together
	family := append(parents, kids...)

	// TODO: Print out the value of `family`
	fmt.Println(family)

	// TODO: Print the length and capacity of the new slice.
	fmt.Printf("len(family): %d, cap(family): %d\n", len(family), cap(family))

	// TODO: Create a new slice called "youngest" by taking the last 3 elements of the slice you just made.
	// HINT: slice[starting_index : (starting_index + length)], or use the shorthand version of slicing

	youngest := family[5:8] // or youngest := family[5:]

	// TODO: Print out the value of the new slice `youngest` that you just created
	fmt.Println(youngest)

	// TODO: Fix the following bugs
	extras := make([]string, 1, 1)
	extras[0] = "Alice"

	fmt.Println(extras)
}
