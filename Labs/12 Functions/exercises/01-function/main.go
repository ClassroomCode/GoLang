package main

import (
	"fmt"
	"strings"
)

func main() {
	// Say "Good morning" to John
	echo(sayHello("John"))
	fmt.Println(Duplicate("gopher"))
	fmt.Println(ProperName("rob", "pike"))
}

// Create a function called `echo` that takes a single argument
// of type string and prints it out.
// func() {}

// Change sayHello to allow the greeting to be customized
func sayHello(name string) string {
	return "Hello " + name
}

// Change the following function to remove the named return
func Duplicate(s string) (dup string) {
	dup = s + " " + s
	return
}

// Change the following function to declare the argument type for each argument provided
func ProperName(first, last string) string {
	return fmt.Sprintf("%s %s", strings.Title(first), strings.Title(last))
}

// Fix the following function declaration
func Broken(s string) string, error {
	return s, errors.New("I'm broken")
}
