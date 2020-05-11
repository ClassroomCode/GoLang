package main

import "fmt"

// section: iota
const (
	Apple int = iota
	Orange
	Banana
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
The value of Orange is 1
The value of Banana is 2
// section: output
*/
