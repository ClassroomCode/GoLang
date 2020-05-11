package main

import (
	"fmt"
	"time"
)

var N = 20

func main() {
	// section: code
	ch := make(chan int)
	for i := 0; i < N; i++ {
		go func(i int) {
			for m := range ch {
				fmt.Printf("routine %d received %d\n", i, m)
			}
		}(i)
	}

	for i := 0; i < N; i++ {
		ch <- i
	}
	time.Sleep(time.Second)
	// section: code
}
