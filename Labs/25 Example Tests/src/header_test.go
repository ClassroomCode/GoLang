package src_test

import (
	"fmt"

	"github.com/gopherguides/learn/_training/testing/example/src"
)

// section: valid
func ExampleHeader() {
	l := src.Header("Title", []string{"Line 1", "Line 2"})
	fmt.Printf(l)
	// Output: <h1>Title</h1>
	//
	// Line 1
	// Line 2
}

// section: valid

// section: bug
func ExampleHeader_bug() {
	l := src.Header("Title", []string{"Line 1", "Line 2"})
	fmt.Printf(l)
	// Output: <h1>Title</h1>

	// Line 1
	// Line 2
}

// section: bug

// section: bug_fixed
func ExampleHeader_bugfixed() {
	l := src.Header("Title", []string{"Line 1", "Line 2"})
	fmt.Printf(l)
	// Output: <h1>Title</h1>
	//
	// Line 1
	// Line 2
}

// section: bugfixed
