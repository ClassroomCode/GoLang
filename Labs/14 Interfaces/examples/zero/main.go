package main

import "fmt"

// section: content
type Foo interface {
	Greet() string
}

func main() {
	var f Foo
	fmt.Println(f.Greet())
}
