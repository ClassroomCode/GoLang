package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const N = 50

var found = make(chan string)

// section: main
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < N; i++ {
		go search(ctx, i)
	}
	m := <-found
	cancel()
	fmt.Print(m)
	time.Sleep(3 * time.Second)
}

// section: main

// section: search
func search(ctx context.Context, i int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Print("x")
			return
		default:
			fmt.Print(".")
			x := rand.Intn(N)
			if x == 42 {
				fmt.Print("found")
				found <- fmt.Sprintf("routine %d: found 42!", i)
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// section: search
