package main

import "fmt"

func main() {
	names := []string{}
	names = append(names, "John")

	// Append multiple items at once
	names = append(names, "Sally", "George")

	// Append an entire slice to another slice
	moreNames := []string{"Bill", "Ginger", "Wilma"}
	names = append(names, moreNames...)

	fmt.Println(names)
}
