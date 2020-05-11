package main

import "fmt"

// section: iota
const (
	Apple int = iota
	Banana
	Orange
)

func main() {
	fmt.Printf("The value of Apple is %v\n", Apple)
	fmt.Printf("The value of Orange is %v\n", Orange)
	fmt.Printf("The value of Banana is %v\n", Banana)
}

// section: iota

/*
// section: output
The value of Apple is 0
The value of Orange is 2
The value of Banana is 1
// section: output
*/
