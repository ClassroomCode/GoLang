package main

import (
	"fmt"
	"log"
)

func main() {
	defer fmt.Println("one")
	log.Fatal("boom")
	defer fmt.Println("three")
}
