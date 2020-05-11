package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	c := NewCommand()

	c.Open()

	sigCh := make(chan os.Signal, 1)

	// Wire up the channel to an `os.Signal`
	signal.Notify(sigCh, os.Interrupt)

	log.Println("Listening for signals")

	<-sigCh
	log.Println("Signal received, initializing clean shutdown...")

	// do graceful shutdown
	// make sure not to block for the timeout added later
	closed := make(chan struct{})
	go func() {
		// do all graceful shutdown of program
		c.Close()

		// signal you've completed your shutdown
		close(closed)
	}()

	log.Println("Waiting for clean shutdown...")

	select {
	case <-sigCh:
		log.Println("second signal received, initializing hard shutdown")
	case <-time.After(time.Second * 30):
		log.Println("time limit reached, initializing hard shutdown")
	case <-closed:
		log.Println("server shutdown completed")
	}

}

func NewCommand() *command {
	wg := sync.WaitGroup{}
	return &command{
		quit: make(chan struct{}),
		wg:   &wg,
	}
}

type command struct {
	quit chan struct{}
	wg   *sync.WaitGroup
}

func (c *command) Open() {
	c.wg.Add(1)
	go c.monitor()
}

func (c *command) Close() {
	// add artificial time to simulate the real world
	time.Sleep(time.Second * 5)
	close(c.quit)
	c.wg.Wait()
}

func (c *command) monitor() {
	defer c.wg.Done()
	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-c.quit:
			log.Println("shutting down monitor")
			return
		case <-ticker.C:
			log.Println("monitor check")
		}
	}
}
