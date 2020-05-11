package main

import "fmt"

type greeter func() string

func main() {
	sayHello(func() string {
		return "Hello"
	})
}

func sayHello(fn greeter) {
	fmt.Println(fn())
}
