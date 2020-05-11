package main

import (
	"fmt"
)

// section: main
func main() {
	c := make(chan int, 10)

	// Write values into a channel
	for i := 0; i < cap(c); i++ {
		c <- i
	}

	// Close the channel
	close(c)

	// You can still read all the data in the channel
	for i := range c {
		fmt.Println(i)
	}

	// but you can no longer write to the channel as it's closed
	c <- 11
}

// section: main
