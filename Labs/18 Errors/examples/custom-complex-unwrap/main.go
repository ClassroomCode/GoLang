package main

import (
	"errors"
	"fmt"
)

// section: code
type ComplexError struct {
	File   string
	Server string
}

func (c ComplexError) Error() string {
	return fmt.Sprintf("%s: %s", c.File, c.Server)
}

func boom() error {
	ce := ComplexError{
		File:   "main.go",
		Server: "myhost",
	}
	err := fmt.Errorf("original error: %w", ce)
	err = fmt.Errorf("wrap again: %w", err)
	return err
}

func main() {
	err := boom()

	// Unwrap and print all errors
	for {
		if err == nil {
			return
		}
		fmt.Println(err)
		err = errors.Unwrap(err)
	}
}

// section: code

/*
// section: output
wrap again: original error: main.go: myhost
original error: main.go: myhost
main.go: myhost
// section: output
*/
