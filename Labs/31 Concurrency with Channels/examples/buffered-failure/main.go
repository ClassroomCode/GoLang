package main

import (
	"fmt"
	"os"
)

func main() {
	// section: code
	messages := make(chan string, 10)

	go func() {
		for i := 0; i < 10; i++ {
			msg := fmt.Sprintf("message %d", i)
			messages <- msg
			fmt.Printf("sent: %s\n", msg)
		}
	}()

	m := <-messages
	fmt.Printf("received: %s\n", m)
	os.Exit(0)
	// sent: message 0
	// sent: message 1
	// sent: message 2
	// received: message 0
	// sent: message 3
	// sent: message 4
	// section: code
}
