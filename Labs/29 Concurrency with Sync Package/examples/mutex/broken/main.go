package main

import (
	"fmt"
	"sync"
)

// section: funcs
var data = map[int]int{}

func get(i int) int {
	return data[i]
}

func set(i int) {
	data[i] = i
}

// section: funcs

// section: main
func main() {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			set(i)
		}(i)
		go func(i int) {
			defer wg.Done()
			fmt.Println(get(i))
		}(i)
	}
	wg.Wait()
}

// section: main
