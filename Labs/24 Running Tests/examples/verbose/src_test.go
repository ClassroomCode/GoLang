package verbose_test

import "testing"


// section: pass
func TestVerbose(t *testing.T) {
	t.Log("This is a log statement for a passing test...")
	// this test passes...

}
// section: pass

/*
// section: pass_output
$ go test -v
=== RUN   TestVerbose
--- PASS: TestVerbose (0.00s)
    verbose_test.go:7: This is a log statement for a passing test...
=== RUN   TestVerboseFail
--- FAIL: TestVerboseFail (0.00s)
    verbose_test.go:25: This is a log statement for a failing test...
    verbose_test.go:26: boom
FAIL
exit status 1
FAIL    github.com/gopherguides/learn/_training/testing/running/src/verbose     0.005s
// section: pass_output
*/

// section: fail
func TestVerboseFail(t *testing.T) {
	t.Log("This is a log statement for a failing test...")
	t.Fatal("boom")
}
// section: fail


/*
// section: fail_output
--- FAIL: TestVerboseFail (0.00s)
    verbose_test.go:32: This is a log statement for a failing test...
    verbose_test.go:33: boom
FAIL
exit status 1
FAIL    github.com/gopherguides/learn/_training/testing/running/src/verbose     0.005s
// section: fail_output
*/


