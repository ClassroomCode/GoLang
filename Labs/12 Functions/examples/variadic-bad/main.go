package main

import "fmt"

func main() {
	sayHello("John", "Paul", "George", "Ringo")
}

func sayHello(names ...string, group string) {
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}
