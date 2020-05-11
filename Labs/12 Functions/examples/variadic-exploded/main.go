package main

import "fmt"

func main() {
	beatles := []string{"John", "Paul", "George", "Ringo"}
	sayHello(beatles...)
	// Hello John
	// Hello Paul
	// Hello George
	// Hello Ringo
}

func sayHello(names ...string) {
	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}
