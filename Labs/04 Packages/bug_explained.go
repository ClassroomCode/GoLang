package main

import "fmt"

func main() {
	fmt.Println("Hello, World")

	// Redeclaring the `fmt` namespace here
	fmt := `my format string: %s`

	// `fmt` is now a string variable, not a package namespace
	fmt.Printf(fmt, "some more text")
}
