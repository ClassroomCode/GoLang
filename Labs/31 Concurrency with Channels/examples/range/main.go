package main

import "fmt"

func main() {
	// empty

}

// section: range
func rangeLoop(m chan string, quit chan struct{}) {
	// loop until the channel is closed
	for m := range messages {
		fmt.Println(m)
	}

	close(quit)
}

// section: range

// section: for
func forLoop(m chan string, quit chan struct{}) {
	// loop until the channel is closed
	for {
		m, ok := <-messages
		if !ok {
			// channel was closed
			break
		}
		fmt.Println(m)
	}
	close(quit)
}

// section: for
