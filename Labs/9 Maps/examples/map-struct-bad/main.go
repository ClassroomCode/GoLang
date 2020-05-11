package main

import "fmt"

// section: code
type simple struct {
	ID   int
	Name string
}

func main() {
	data := map[int]simple{}
	s1 := simple{ID: 1, Name: "Cory"}
	data[1] = s1

	value := data[10]

	// Because of the way Zero values in Go work,
	// we still get an `zero` value representation of the struct
	// which is certainly a bug in production
	fmt.Printf("%+v", value)
}

// section: code

/*

// section: expected
{ID:1 Name:Cory}
// section: expected

// section: output
{ID:0 Name:}
// section: output

*/
