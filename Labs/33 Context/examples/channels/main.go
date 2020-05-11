package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// section: chan
const N = 50

var found = make(chan string)
var quit = make(chan struct{})

func main() {
	for i := 0; i < N; i++ {
		go search(i)
	}
	fmt.Print(<-found)
	close(quit)
	// section: chan
	time.Sleep(3 * time.Second)
}

// section: search
func search(i int) {
	for {
		select {
		case <-quit:
			fmt.Print("x")
			return
		default:
			fmt.Print(".")
			x := rand.Intn(N)
			if x == 42 {
				fmt.Print("found")
				found <- fmt.Sprintf("routine %d: found 42!", i)
				break
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// section: search
