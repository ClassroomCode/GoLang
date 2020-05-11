package src

import "testing"

// section: test
func TestTableDriven(t *testing.T) {
	tcs := []struct {
		name     string
		A        int
		B        int
		Expected int
	}{
		{name: "1-1", A: 1, B: 1, Expected: 2},
		{name: "2-2", A: 2, B: 2, Expected: 4},
		{name: "2-2", A: 3, B: 3, Expected: 5},
		{name: "2-2", A: 4, B: 4, Expected: 6},
	}

	globalSetup()          // run necessary setup for all the tests
	defer globalTeardown() // make sure tearing down of all the tests

	for _, tc := range tcs {
		tc := tc // rebind lexical scope
		t.Run(tc.name, func(t *testing.T) {
			setup() // run necessary setup for the test

			got := tc.A + tc.B
			if got != tc.Expected {
				t.Errorf("expected %d, got %d", tc.Expected, got)
			}

			teardown() // run necessary teardown of the test
		})
	}
}

// section: test

func globalSetup()    {}
func globalTeardown() {}
func setup()          {}
func teardown()       {}

/*
// section: test-output
--- FAIL: TestTableDriven (0.00s)
 table_driven_test.go:20: expected 5, got 6
 table_driven_test.go:20: expected 6, got 8
// section: test-output
*/
