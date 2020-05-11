package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

type Persons []Person

// Len is the number of elements in the collection.
func (p Persons) Len() int {
	return len(p)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (p Persons) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

// Swap swaps the elements with indexes i and j.
func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	var people = Persons{
		{
			Name: "Bill",
			Age:  55,
		},
		{
			Name: "John",
			Age:  92,
		},
		{
			Name: "Megan",
			Age:  4,
		},
	}
	fmt.Printf("unsorted: %+v\n", people)
	sort.Sort(people)
	fmt.Printf("sorted: %+v\n", people)
}
