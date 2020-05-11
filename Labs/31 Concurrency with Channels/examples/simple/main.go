package main

import "fmt"

func main() {

	// make a new channel of type string
	messages := make(chan string)

	// this goroutine launches immediately
	go func() {

		// This line blocks until the channel is read from
		fmt.Println("waiting to write to messages")
		messages <- "hello!"
		fmt.Println("wrote to messages")
	}()

	// this line blocks until someone writes to the channel
	fmt.Println("waiting to read from messages")
	msg := <-messages
	fmt.Println("read from messages")

	fmt.Println(msg)
}
