package main

import "fmt"

func main() {
	fmt.Println("Hello, World")

	fmt := `my format string: %s`

	fmt.Printf(fmt, "some more text")
}
