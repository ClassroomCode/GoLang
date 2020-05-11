package src

import "testing"

// section: test
func TestSimple(t *testing.T) {
	if true {
		t.Error("expected false, got true")
	}
}

// section: test

/* output
// section: output
--- FAIL: TestSimple (0.00s)
 simple_test.go:7: expected false, got true
FAIL
// section: output
*/
