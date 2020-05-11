package main

import "fmt"

// Create a function called `printSlice` that takes a slice of strings
// and prints them out.
// func () {}

// Create a function called `printAll` that takes a string as a variadic
// argument and prints them out.
// func () {}

func main() {
	fruit := []string{"banana", "orange", "strawberry", "apple"}
	fmt.Println("Print out a slice")
	printSlice(fruit)
	fmt.Println()

	fmt.Println("Print out multiple items")
	printAll("banana", "orange", "strawberry", "apple")

	// Explode the arguments to pass a slice to printAll
	fmt.Println() // insert blank line in output
	printAll(fruit...)

	// You can also call printAll with no arguments
	fmt.Println() // insert blank line in output
	printAll()
}
