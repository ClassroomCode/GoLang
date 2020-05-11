package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"training/cmd"
)

// section: main
func main() {

	c := cmd.New()
	if err := c.Open(); err != nil {
		log.Fatal(c)
	}

	sigCh := make(chan os.Signal, 1)

	// Wire up the channel to an `os.Signal`
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	log.Println("Listening for signals")

	// launch all program code...

	// Block until one of the signals above is received
	<-sigCh
	log.Println("Signal received, initializing clean shutdown...")

	// do graceful shutdown
	// make sure not to block for the timeout added later
	closed := make(chan struct{})
	go func() {
		// do all graceful shutdown of program
		if err := c.Close(); err != nil {
			log.Fatal(err)
		}

		// signal you've completed your shutdown
		close(closed)
	}()

	// Block again until another signal is received, a shutdown timeout elapses,
	// or the Command is gracefully closed
	log.Println("Waiting for clean shutdown...")
	select {
	case <-sigCh:
		log.Println("second signal received, initializing hard shutdown")
	case <-time.After(time.Second * 30):
		log.Println("time limit reached, initializing hard shutdown")
	case <-closed:
		log.Println("server shutdown completed")
	}

	// goodbye.
}

// section: main

/*
// section: output
2020/04/20 11:22:35 Listening for signals
2020/04/20 11:22:36 monitor check
2020/04/20 11:22:37 monitor check
^C2020/04/20 11:22:37 Signal received, initializing clean shutdown...
2020/04/20 11:22:37 Waiting for clean shutdown...
2020/04/20 11:22:37 shutting down monitor
2020/04/20 11:22:42 server shutdown completed
// section: output
*/
