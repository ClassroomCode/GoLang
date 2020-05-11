package main

import (
	"fmt"
)

func main() {
	// section: main
	// Create a new pointer to a string
	ps := new(string)

	// Set the value of ps by dereferencing the pointer
	*ps = "hello"

	// Capture a new "value" as a string
	s := *ps

	// section: main
	fmt.Printf("ps type:\t\t%T\nps value:\t\t%v\ndereferenced value:\t%q\n", ps, ps, *ps)
	fmt.Printf("s type:\t\t\t%T\ns value:\t\t%v\n", s, s)
}

/*
// section: output
ps type:                *string
ps value:               0xc000092030
dereferenced value:     "hello"
s type:                 string
s value:                hello
// section: output
*/
