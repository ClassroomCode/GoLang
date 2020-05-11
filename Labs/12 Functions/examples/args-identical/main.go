package main

import "fmt"

func main() {
	sayHello("Hello", "Mark")
	sayHello2("Hello", "Mark")
}

func sayHello(greeting, name string) {
	fmt.Printf("%s, %s", greeting, name)
}

func sayHello2(greeting string, name string) {
	fmt.Printf("%s, %s", greeting, name)
}
