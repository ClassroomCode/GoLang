package src_test

import "testing"

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

	// section: parallel
	// iterate over all the tests
	for _, tc := range tcs {
		tc := tc // rebind tc into this lexical scope
		// start a subtest
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			// use the test values from tc to actually run the test
		})
	}
	// section: parallel
}
