package main

import (
	"fmt"
)

func process(messages chan string, quit chan struct{}) {
	// loop until the channel is closed
	for m := range messages {
		fmt.Println(m)
	}

	close(quit)
}

func main() {
	messages := make(chan string, 5)
	quit := make(chan struct{})

	go process(messages, quit)

	fruits := []string{"apple", "plum", "peach", "pear", "grape"}
	for _, s := range fruits {
		messages <- s
	}
	// closing the messages channel will end the for loop
	// in the processes goroutine
	close(messages)

	<-quit
	fmt.Println("done")
}
