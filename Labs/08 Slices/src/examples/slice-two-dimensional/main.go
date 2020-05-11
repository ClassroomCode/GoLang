package main

import "fmt"

func main() {
	// a slice of string slices
	// section: slice
	type Modules [][]string

	course := Modules{
		[]string{"Chapter One: Syntax"},
		[]string{"Chapter Two: Arrays and Slices"},
		[]string{"Chapter Three: Maps", "Chapter Four: Packages"},
	}
	// section: slice
	fmt.Println(course)
}
