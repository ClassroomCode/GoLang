package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {

	go process()

	// wait to exit based on user input
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

// section: process
func process() {
	// create a timeout.
	timeout := time.NewTimer(1 * time.Second)
	// this ensures we don't leak a go routine
	defer timeout.Stop()

	// create a channel to send our result
	res := make(chan int, 1)

	// create another go routine to do some work
	go func() {
		time.Sleep(2 * time.Second)
		res <- 1
		fmt.Println("wrote to result channel")
	}()

	select {
	case <-res:
		fmt.Println("sub routine executed first")
	case <-timeout.C:
		fmt.Println("timeout executed first")
	}
}

// section: process
