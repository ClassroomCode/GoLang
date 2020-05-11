package main

import "fmt"

func main() {
	defer goodbye()
	fmt.Println("hello")
}

func goodbye() {
	fmt.Println("goodbye")
}
