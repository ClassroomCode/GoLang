package main

import (
	"log"
	"os"
	"os/signal"
	"training/cmd"
)

// section: main
func main() {

	c := cmd.New()
	if err := c.Open(); err != nil {
		log.Fatal(c)
	}

	quit := make(chan os.Signal, 1)

	// Wire up the channel to an `os.Signal`
	signal.Notify(quit, os.Interrupt)

	log.Println("program started")

	// Block until a signal is received.
	s := <-quit

	log.Printf("received %s signal for shutdown", s)

	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("successfuly shutdown")
}

// section: main

/*
// section: output
2020/04/20 11:17:03 program started
2020/04/20 11:17:04 monitor check
2020/04/20 11:17:05 monitor check
2020/04/20 11:17:06 monitor check
^C2020/04/20 11:17:06 received interrupt signal for shutdown
2020/04/20 11:17:06 shutting down monitor
2020/04/20 11:17:06 successfuly shutdown
// section: output
*/
