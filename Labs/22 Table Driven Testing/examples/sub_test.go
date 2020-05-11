package src

import (
	"fmt"
	"testing"
)

// section: test
func TestSub(t *testing.T) {
	tt := []struct {
		A        int
		B        int
		Expected int
	}{
		{A: 1, B: 1, Expected: 2},
		{A: 2, B: 2, Expected: 4},
		{A: 3, B: 3, Expected: 5},
		{A: 4, B: 4, Expected: 6},
	}

	globalSetup()          // run necessary setup for all the tests
	defer globalTeardown() // make sure tearing down of all the tests

	for i, x := range tt {
		t.Run(fmt.Sprintf("sub test (%d)", i), func(st *testing.T) {
			setup()          // run necessary setup for the test
			defer teardown() // make sure teardown is always called

			got := x.A + x.B
			if got != x.Expected {
				st.Errorf("expected %d, got %d", x.Expected, got)
			}
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
--- FAIL: TestSub (0.00s)
    --- FAIL: TestSub/sub_test_(2) (0.00s)
     sub_test.go:24: expected 5, got 6
    --- FAIL: TestSub/sub_test_(3) (0.00s)
     sub_test.go:24: expected 6, got 8

// section: test-output
*/
