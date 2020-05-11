package main

import "fmt"

// section: code
type prettyString string

func (p prettyString) String() string {
	return fmt.Sprintf("I'm a pretty string: %q", string(p))
}

func main() {
	s := "Hi"
	fmt.Println(s)
	fmt.Println(prettyString(s))
}

// section: code

/*
// section: output
Hi
I'm a pretty string: "Hi"
// section: output
*/
