package main

import (
	"context"
	"fmt"
	"log"
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctx = context.WithValue(ctx, "magic", 42)
	for i := 0; i < N; i++ {
		ctx = context.WithValue(ctx, "id", i)
		go search(ctx)
	}
	select {
	case m := <-found:
		cancel()
		fmt.Print(m)
	case <-ctx.Done():
		cancel()
		if err := ctx.Err(); err != nil {
			fmt.Println(err)
		}
	}
	time.Sleep(3 * time.Second)
}

// section: main

// section: search
func search(ctx context.Context) {
	magic, ok := ctx.Value("magic").(int)
	if !ok {
		log.Fatalf("was expecting 'magic' to be an int, but got %T", ctx.Value("magic"))
	}
	i, ok := ctx.Value("id").(int)
	if !ok {
		log.Fatalf("was expecting 'id' to be an int, but got %T", ctx.Value("id"))
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Print("x")
			return
		default:
			fmt.Print(".")
			x := rand.Intn(N)
			if x == magic {
				fmt.Print("found")
				found <- fmt.Sprintf("routine %d: found %d!", i, magic)
				return
			}
			time.Sleep(10 * time.Millisecond)
		}
	}
}

// section: search
