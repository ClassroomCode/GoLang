package main

import (
	"errors"
	"fmt"
)

// section: main
type errFound string

func (e *errFound) Error() string {
	return "found"
}

func main() {
	err := exists("README.md")
	if err != nil {
		fmt.Println("error was NOT nil")
		var ef *errFound
		if errors.As(err, &ef) {
			fmt.Println("this was a errFound...")
		}
	}
}

func exists(path string) error {
	if path == "README.md" {
		// return a `typed` nil value...
		var e *errFound
		if e == nil {
			fmt.Println("yup, it's nil here..")
		}
		return e
	}
	return nil
}

// section: main

/*
// section: output
yup, it's nil here..
error was NOT nil
this was a errFound...
// section: output
*/
