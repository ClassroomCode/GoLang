package main

import "fmt"

func main() {
	defer fmt.Println("one")
	defer panic("two")
	defer fmt.Println("three")
}
