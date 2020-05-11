package main

import "fmt"

func main() {
	// make a full directional channel
	c := make(chan string)

	go send(c)

	// read from the channel
	for s := range c {
		fmt.Println(s)
	}

}

// send takes the first argument as a send only channel
func send(c chan<- string) {
	// Write into the channel
	for _, s := range []string{"one", "two", "three"} {
		c <- s
	}
	// You can't read from this channel:
	//s := <-c

	// it is ok to close this channel
	close(c)
}
