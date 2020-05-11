package main

import "fmt"

// section: interfaces
type Greeter interface {
	Greet(string)
}

type Welcomer interface {
	Greet(string)
}

// section: interfaces

// section: person
type Person struct{}

func (p *Person) Greet(name string) {
	fmt.Println("Hello", name)
}

// section: person

func main() {
	// section: assign
	p := &Person{}
	var g Greeter = p
	var w Welcomer = g

	g.Greet("world")
	w.Greet("tim")
	// section: assign
}
