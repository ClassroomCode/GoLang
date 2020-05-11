package main

import (
	"errors"
	"fmt"
)

// section: main
var errUnknown = errors.New("unknown error")

type ComplexError struct {
	File   string
	Server string
	err    error
}

func (c ComplexError) Error() string {
	return fmt.Sprintf("%s: %s", c.File, c.Server)
}

func (c ComplexError) Unwrap() error {
	return c.err
}

func main() {
	err := boom()

	if errors.Is(err, errUnknown) {
		fmt.Println("found simple error")
	}

	var ce ComplexError
	if errors.As(err, &ce) {
		fmt.Println("found complex error")
	}
}

func boom() error {
	ce := ComplexError{
		File:   "README.md",
		Server: "localhost:9090",
		err:    errUnknown,
	}
	return fmt.Errorf("wrapping a complex error: %w", ce)
}

// section: main

/*
// section: output
found simple error
found complex error
// section: output
*/
