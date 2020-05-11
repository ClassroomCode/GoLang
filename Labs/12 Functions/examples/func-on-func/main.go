package main

import "fmt"

// section: greeter
type greeter func() string

func (greeter) Name() string {
	return "World"
}

func sayHello(fn greeter) {
	fmt.Println(fn(), fn.Name())
}

// section: greeter

func main() {
	sayHello(func() string {
		return "Hello"
	})
}
