package main

import "fmt"

func main() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
}
