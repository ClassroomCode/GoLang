package src_test

import "testing"

// section: test
func TestSomething(t *testing.T) {
	tcs := []struct {
		name  string // the name of the subtest
		input string // arbitrary inputs
		exp   string // what you expect to see
	}{
		{
			name:  "some test case",
			input: "some input",
			exp:   "some expected output",
		},
	}

	// iterate over all the tests
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			// use the test values from tc to actually run the test
		})
	}
}

// section: test
