package src

import (
	"errors"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// section: helpers

// section: ok
// ok fails the test if an err is not nil.
func ok(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

// section: ok

// section: equals
// equals fails the test if exp is not equal to got.
func equals(t testing.TB, got, exp interface{}) {
	t.Helper()
	if !cmp.Equal(exp, got) {
		t.Fatalf("\n\tgot: %#v\n\texp: %#v\n", got, exp)
	}
}

// section: equals

// section: helpers

func DoSomething() error {
	return errors.New("failed to process")
}

// section: ok-test
func TestOk(t *testing.T) {
	err := DoSomething()
	ok(t, err)
}

// section: ok-test

/*
// section: ok-output
$ go test ./helpers_test.go  -run TestOk
--- FAIL: TestOk (0.00s)
    helpers_test.go:43: unexpected error: failed to process
FAIL
FAIL    command-line-arguments  0.006s
// section: ok-output
*/

// section: equals-test
func TestEquals(t *testing.T) {
	v := strings.Title("foo")
	equals(t, v, "foo")
}

// section: equals-test

/*
// section: equals-output
$ go test ./helpers_test.go  -run TestEquals
--- FAIL: TestEquals (0.00s)
    helpers_test.go:28:
                exp: "Foo"
                got: "foo"
FAIL
FAIL    command-line-arguments  0.007s
// section: equals-output
*/
