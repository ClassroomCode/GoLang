package main

import (
	"fmt"
	"os"
)

// section: main
func main() {
	c := make(chan string)
	quit := make(chan struct{})

	go func(messages []string) {
		for _, s := range messages {
			c <- s
		}
		close(quit)
	}([]string{"hi", "bye"})

	for {
		select {
		case message := <-c:
			fmt.Println(message)
		case <-quit:
			fmt.Println("shutting down")
			os.Exit(0)
		}
	}
}

// section: main
