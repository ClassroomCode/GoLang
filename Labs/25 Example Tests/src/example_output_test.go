package src_test

import (
	"fmt"

	"github.com/gopherguides/learn/_training/fundamentals/testing/src"
)

// This is some text from your example.
// You can add additional information about your example here.
func ExampleDuplicate_output() {
	d := src.Duplicate("foo")
	fmt.Println(d)
	// Output: foofoo
}
