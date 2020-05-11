package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	// Say "Good morning" to John
	echo(sayHello("Good morning", "John"))
	fmt.Println(Duplicate("gopher"))
	fmt.Println(ProperName("rob", "pike"))
}

// Create a function called `echo` that takes a single argument
// of type string and prints it out.
func echo(s string) {
	fmt.Println(s)
}

// Change sayHello to allow the greeting to be customized
func sayHello(greeting string, name string) string {
	return greeting + " " + name
}

// Change the following function to remove the named return
func Duplicate(s string) string {
	return s + " " + s
}

// Change the following function to declare the argument type for each argument provided
func ProperName(first string, last string) string {
	return fmt.Sprintf("%s %s", strings.Title(first), strings.Title(last))
}

// Fix the following function declaration
func Broken(s string) (string, error) {
	return s, errors.New("I'm broken")
}
