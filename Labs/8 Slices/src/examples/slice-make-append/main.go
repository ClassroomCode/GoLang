package main

import "fmt"

func main() {
	a := make([]string, 2)
	a = append(a, "foo", "bar")
	fmt.Printf("%q", a)
}
