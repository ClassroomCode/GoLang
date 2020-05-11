package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)     // add 1 to the internal counter
	go sayHello() // schedule this to run in a goroutine
	wg.Wait()     // block here until we decrement the internal counter to 0
}

func sayHello() {
	fmt.Println("hello")
	wg.Done() // decrement the internal counter
}
