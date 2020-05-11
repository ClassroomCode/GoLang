package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("one")
	os.Exit(1)
	defer fmt.Println("three")
}
