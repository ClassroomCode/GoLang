package main

import "fmt"

func main() {
	// create an array of names
	namesArray := [4]string{"John", "Paul", "George", "Ringo"}
	// create a slice of names
	namesSlice := []string{"John", "Paul", "George", "Ringo"}

	fmt.Println(namesArray)
	fmt.Println(namesSlice)
}
