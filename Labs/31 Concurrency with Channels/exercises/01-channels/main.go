package main

import "fmt"

func process(messages chan string, quit chan struct{}) {
	// Using the `range` operator, loop through messages
	// until the channel is closed
	for m := range messages {
		// print the message with a for loop using range
	}

	// close the `quit` channel
}

func main() {
	// declare the messages channel of type string and capacity of 5

	// declare a signal channel

	// launch the `process` function in a goroutine

	// declare 5 fruits in a []string

	// loop through the fruits and send them to the messages channel
	for _, f := range fruits {
	}

	// close the messages channel

	// wait for everything to finish (hint, block on the quit channel)

	fmt.Println("done")
}
