package main

import (
	"errors"
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
		log.Println("error: ", err)
	}
	var ce ComplexError
	if errors.As(err, &ce) {
		log.Println("server:", ce.Server)
	}
}

// section: main
