package main

import (
	"fmt"
	"os"
	"os/signal"
)

// section: main
func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)

	// Wire up the channel to an `os.Signal`
	signal.Notify(c, os.Interrupt)

	fmt.Println("awaiting signal...")

	// Block until a signal is received.
	s := <-c

	fmt.Println("Got signal:", s)
	// do my code shutdown....
	// then exit
}

// section: main
