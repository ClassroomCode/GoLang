package main

import (
	"fmt"
	"sync"
)

// section: main
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	var mu sync.Mutex

	var count int

	go func() {
		for i := 0; i < 5; i++ {
			mu.Lock()
			count = i
			fmt.Println(count)
			mu.Unlock()
		}
		wg.Done()
	}()

	wg.Wait()
}

// section: main
