package main

import "fmt"

const (
	_   int = iota // Skip the first value of 0
	Foo            // Foo = 1
	Bar            // Bar = 2
	_
	_
	Bin // Bin = 5
	// Using a comment or a blank line will not increment the iota value

	Baz // Baz = 6
)

func main() {
	fmt.Printf("The value of Foo is %v\n", Foo)
	fmt.Printf("The value of Bar is %v\n", Bar)
	fmt.Printf("The value of Bin is %v\n", Bin)
	fmt.Printf("The value of Baz is %v\n", Baz)
}
