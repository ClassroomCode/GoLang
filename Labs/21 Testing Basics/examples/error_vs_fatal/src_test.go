package src

import (
	"errors"
	"fmt"
	"testing"
)

// section: test
func TestNewObject(t *testing.T) {
	obj, err := newObject()
	if err != nil {
		// We need to fatal and end the tests, as `obj` is now `nil` and future
		// tests would panic or fail
		t.Fatal(err)
	}
	// continue with more tests
	fmt.Println(obj)
}

// section: test

// supporting code for above test

type object struct{}

func newObject() (*object, error) {
	return nil, errors.New("error creating object")
}
