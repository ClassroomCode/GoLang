package main

import "fmt"

func main() {
	f := func() string {
		return "Hello"
	}
	sayHello(f)
}

func sayHello(fn func() string) {
	fmt.Println(fn())
}
