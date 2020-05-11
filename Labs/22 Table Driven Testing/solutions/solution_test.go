package src_test

import (
	"strings"
	"testing"
)

// section: solution
func Test_Index(t *testing.T) {
	tcs := []struct {
		value  string
		substr string
		exp    int
	}{
		{"Gophers are amazing", "are", 8},
		{"Testing in Go is fun.", "fun", 17},
		{"The answer is 42.", "is", 11},
	}
	for _, tc := range tcs {
		got := strings.Index(tc.value, tc.substr)
		if got != tc.exp {
			t.Errorf("unexpected value.  got %d, exp %d", got, tc.exp)
		}
	}
}

// section: solution

// section: solutionsub
func Test_Index_Subtest(t *testing.T) {
	tcs := []struct {
		value  string
		substr string
		exp    int
	}{
		{"Gophers are amazing", "are", 8},
		{"Testing in Go is fun.", "fun", 17},
		{"The answer is 42.", "is", 11},
	}
	for _, tc := range tcs {
		t.Run(tc.value, func(t *testing.T) {
			tc := tc // rebind the lexical scope
			t.Parallel()
			got := strings.Index(tc.value, tc.substr)
			if got != tc.exp {
				t.Errorf("unexpected value.  got %d, exp %d", got, tc.exp)
			}
		})
	}
}

// section: solutionsub

/*
// section: output
$ go test -v ./solution_test.go
=== RUN   Test_HasPrefix
--- PASS: Test_HasPrefix (0.00s)
=== RUN   Test_Index
--- PASS: Test_Index (0.00s)
=== RUN   Test_Index_Subtest
=== RUN   Test_Index_Subtest/Gophers_are_amazing
=== PAUSE Test_Index_Subtest/Gophers_are_amazing
=== RUN   Test_Index_Subtest/Testing_in_Go_is_fun.
=== PAUSE Test_Index_Subtest/Testing_in_Go_is_fun.
=== RUN   Test_Index_Subtest/The_answer_is_42.
=== PAUSE Test_Index_Subtest/The_answer_is_42.
=== CONT  Test_Index_Subtest/Gophers_are_amazing
=== CONT  Test_Index_Subtest/The_answer_is_42.
=== CONT  Test_Index_Subtest/Testing_in_Go_is_fun.
--- PASS: Test_Index_Subtest (0.00s)
    --- PASS: Test_Index_Subtest/Gophers_are_amazing (0.00s)
    --- PASS: Test_Index_Subtest/Testing_in_Go_is_fun. (0.00s)
    --- PASS: Test_Index_Subtest/The_answer_is_42. (0.00s)
=== RUN   Test_HasSuffix
--- PASS: Test_HasSuffix (0.00s)
PASS
ok      command-line-arguments  0.006s

// section: output
*/
