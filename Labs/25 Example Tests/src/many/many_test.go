package many_test

import (
	"fmt"

	"github.com/gopherguides/learn/_training/testing/example/src/many"
)

// section: examples
func ExampleFind() {
	values := []string{"Apple", "Orange", "Pear", "Apple", "Watermelon"}
	r := many.Find("Apple", values, many.All)
	fmt.Println(r)
	// Output: [0 3]
}

func ExampleFind_first() {
	values := []string{"Apple", "Orange", "Pear", "Apple", "Watermelon"}
	r := many.Find("Pear", values, many.First)
	fmt.Println(r)
	// Output: [2]
}

func ExampleFind_last() {
	values := []string{"Apple", "Orange", "Pear", "Apple", "Watermelon"}
	r := many.Find("Apple", values, many.Last)
	fmt.Println(r)
	// Output: [3]
}

// section: examples
