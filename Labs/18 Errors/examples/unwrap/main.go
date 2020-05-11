package main

import (
	"fmt"
	"io"

	"errors"
)

// section: code
type ComplexError struct {
	File   string
	Server string
	err    error
}

func (c ComplexError) Error() string {
	return fmt.Sprintf("%s: %s - %s", c.File, c.Server, c.err)
}

func (c ComplexError) Unwrap() error {
	return c.err
}

func boom() error {
	return ComplexError{
		File:   "main.go",
		Server: "myhost",
		err:    io.EOF,
	}
}

func main() {
	err := boom()
	if e := errors.Unwrap(err); err != nil {
		fmt.Printf("original error was: %s", e)
	}
}

// section: code

/*
// section: output
original error was: EOF
// section: output
*/
