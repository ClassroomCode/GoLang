package main

import "fmt"

func main() {
	sayHello("John", "Paul", "George", "Ringo")
	// Hello John
	// Hello Paul
	// Hello George
	// Hello Ringo
	sayHello("George")
	// George
	sayHello()
	//
}

func sayHello(names ...string) {
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}
