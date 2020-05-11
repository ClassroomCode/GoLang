package main

import "fmt"

type greeter func() string

func (g greeter) name() string {
	return "Ringo"
}

func sayHello(fn greeter) {
	fmt.Printf("%s %s", fn(), fn.name()) // Hello Ringo
}

func main() {
	sayHello(func() string {
		return "Hello"
	})
}
