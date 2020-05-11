package main

import (
	"fmt"
)

func main() {
	takeTwo(returnTwo())
}

func returnTwo() (string, string) {
	return "hello", "world"
}

func takeTwo(a, b string) {
	fmt.Println(a, b)
}
