package main

import "fmt"

func main() {
	// Adding a second argument to the make function creates a buffered channel
	messages := make(chan string, 2)

	// The program is no longer blocked on writing to a channel,
	// as it has capacity to write
	messages <- "hello!"
	messages <- "hello again!"

	// Reads are no longer blocked as there is already something to read from
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
