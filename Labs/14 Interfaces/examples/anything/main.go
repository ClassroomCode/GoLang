package main

import (
	"fmt"
)

func main() {
	print(42)
	print("Ringo Rules")
	print(3.14159)
	// int: 42
	// string: Ringo Rules
	// float64: 3.14159
}

func print(i interface{}) {
	// %T - prints the type of the value
	// %v - prints the value in a default format
	fmt.Printf("%T: %v\n", i, i)
}
