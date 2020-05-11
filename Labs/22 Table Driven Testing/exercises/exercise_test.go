package src_test

import (
	"strings"
	"testing"
)

// Write a table drive test for strings.Index
// https://golang.org/pkg/strings/#Index
// Use the following test conditions
// |------------------------------------------------|
// | Value                     | Substring | Answer |
// |===========================|===========|========|
// | "Gophers are amazing!"    | "are"     | 8      |
// | "Testing in Go is fun."   | "fun"     | 17     |
// | "The answer is 42."       | "is"      | 11     |
// |------------------------------------------------|

// Rewrite the above test for strings.Index using subtests
// Enable parallel testing in the subtests

// Here is a simple test that tests `strings.HasSuffix`.
// https://golang.org/pkg/strings/#HasSuffix
func Test_HasSuffix(t *testing.T) {
	value := "main.go"
	if !strings.HasSuffix(value, ".go") {
		t.Fatalf("expected %s to have suffix %s", value, ".go")
	}
}
