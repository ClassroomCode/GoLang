package main

import (
	"fmt"
	"sync"
)

// section: main
func main() {
	var wg sync.WaitGroup

	var mu sync.Mutex

	var count int

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			count = i
			fmt.Println(count)
			wg.Done()
		}(i)
	}

	wg.Wait()
}

// section: main
