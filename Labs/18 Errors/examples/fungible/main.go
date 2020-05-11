package main

import (
	"errors"
	"fmt"
)

// section: customerror
type Error string

func (e Error) Error() string {
	return string(e)
}

// section: customerror

func main() {

	// section: const
	const str1 = "EOF"
	const str2 = "EOF"

	fmt.Printf("str1 == str2: %v\n", str1 == str2) // true
	// section: const

	// section: consterr
	const err1 = Error("EOF")
	const err2 = Error("EOF")

	fmt.Printf("err1 == err2: %v\n", err1 == err2) // true
	// section: consterr

	// section: varerr
	var err3 = errors.New("EOF")
	var err4 = errors.New("EOF")

	fmt.Printf("err3 == err4: %v\n", err3 == err4) // false
	// section: varerr
}
