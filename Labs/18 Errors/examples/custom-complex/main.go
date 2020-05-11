package main

import (
	"fmt"
	"log"
)

// section: error
type ComplexError struct {
	File   string
	Server string
}

func (c ComplexError) Error() string {
	return fmt.Sprintf("%s: %s", c.File, c.Server)
}

func boom() error {
	return ComplexError{
		File:   "main.go",
		Server: "myhost",
	}
}

// section: error

// section: main
func main() {
	err := boom()
	if err != nil {
		log.Println(err)
	}
	if ce, ok := err.(ComplexError); ok {
		log.Println(ce.Server)
	}
}

// section: main
