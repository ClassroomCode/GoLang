package main

import "fmt"

func main() {
	// make a full directional channel
	c := make(chan string)
	quit := make(chan struct{})

	go receive(c, quit)

	// Write into the channel
	for _, s := range []string{"one", "two", "three"} {
		c <- s
	}
	close(c)
	// wait for quit channel to close
	<-quit
}

// receive takes the first argument as a read only channel
func receive(c <-chan string, quit chan struct{}) {
	// read from the channel
	for s := range c {
		fmt.Println(s)
	}
	// You can't write to this channel:
	//c <- "foo"

	// You also can't close a read only channel.  This is to protect the sender from you also polluting this channel:
	//close(c)

	close(quit)
}
