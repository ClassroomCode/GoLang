package main

import "fmt"

func main() {
	x := "George"
	printValue(&x)
	fmt.Println(x)

	// using short variable declaration, and the `new` keyword,
	// create and initialize a variable named `s` to a pointer to a string

	// set the value of `s` to "hello"

	// print out both the pointer value and the dereferenced value of `s`
}

func printValue(s *string) {
	// print the pointer value
	// print the string value
	// change the string value
}
