package main

import (
	"fmt"
	"sync"
)

func main() {
	// section: main
	var wg sync.WaitGroup

	names := []string{"Fred", "Wilma", "Barney", "Betty"}

	// Add the len of names as we will be launching a Go routine for each one
	wg.Add(len(names))
	for i, n := range names {
		// launch each process in a goroutine
		go func() {
			fmt.Printf("%q is %d at index position\n", n, i)
			wg.Done() // signal that this Go routine has finished
		}()
	}
	wg.Wait() // Wait for all goroutines to finish
	// section: main
}
