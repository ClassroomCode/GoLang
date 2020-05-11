package main

import "fmt"

func main() {
	name := "Ringo"
	f := func() {
		fmt.Printf("Hello %s", name)
	}

	f() // Hello Ringo
}
