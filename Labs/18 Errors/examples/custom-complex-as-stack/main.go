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

// section: wrap
func boom() error {
	ce := ComplexError{
		File:   "main.go",
		Server: "myhost",
	}
	err := fmt.Errorf("original error: %w", ce)
	err = fmt.Errorf("wrap again: %w", err)
	return err
}

// section: wrap

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

/*
// section: output
2020/04/18 12:47:25 error:  wrap again: original error: main.go: myhost
2020/04/18 12:47:25 server: myhost
// section: output
*/
