package main

import (
	"fmt"
)

func main() {
	b, err := someFunction(1)
	fmt.Println("### b ->", b)
	fmt.Println("### err ->", err)

	b, err = someFunction(-1)
	fmt.Println("### b ->", b)
	fmt.Println("### err ->", err)
}

// section: named
func someFunction(val int) (ok bool, err error) {
	if val == 0 {
		return false, nil
	}
	if val < 0 {
		return false, fmt.Errorf("value can't be negative %d", val)
	}
	ok = true
	return
}

// section: named
