package math_test

import (
	"testing"

	"github.com/gopherguides/learn/_training/fundamentals/testing/src/math"
)

// section: bad
func TestAddTen(t *testing.T) {
	v := math.AddTen(1)
	if v != 11 {
		t.Fatal("unexpected value")
	}
}

// section: bad

/*
// section: bad-output
--- FAIL: TestAddTen (0.00s)
	math_test.go:13: unexpected value
// section: bad-output
*/

// section: better
func TestAddTen_Better(t *testing.T) {
	v := math.AddTen(1)
	if v != 11 {
		t.Fatalf("unexpected value, got: %d", v)
	}
}

// section: better

/*
// section: better-output
--- FAIL: TestAddTen_Better (0.00s)
  math_test.go:33: unexpected value, got: 12
// section: better-output
*/

// section: better-yet
func TestAddTen_BetterYet(t *testing.T) {
	v := math.AddTen(1)
	if v != 11 {
		t.Fatalf("unexpected value, got: %d, exp: %d", v, 11)
	}
}

// section: better-yet

/*
// section: better-yet-output
--- FAIL: TestAddTen_BetterYet (0.00s)
  math_test.go:33: unexpected value, got: 12, exp: 11
// section: better-yet-output
*/

// section: best
func TestAddTen_Best(t *testing.T) {
	got := math.AddTen(1)
	exp := 11
	if got != exp {
		t.Fatalf("unexpected value, got: %d, exp: %d", got, exp)
	}
}

// section: best

/*
// section: best-output
--- FAIL: TestAddTen_Best (0.00s)
  math_test.go:42: unexpected value, got: 12, exp: 11
// section: best-output
*/

// section: best-condensed
func TestAddTen_Best(t *testing.T) {
	if got, exp := math.AddTen(1), 11; got != exp {
		t.Fatalf("unexpected value, got: %d, exp: %d", got, exp)
	}
}

// section: best-condensed
