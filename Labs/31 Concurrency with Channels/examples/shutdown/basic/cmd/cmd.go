package cmd

import (
	"log"
	"time"
)

func New() *Command {
	return &Command{
		quit: make(chan struct{}),
	}
}

type Command struct {
	quit chan struct{}
}

func (c *Command) Open() error {
	go c.monitor()

	return nil
}

func (c *Command) Close() error {
	close(c.quit)
	return nil
}

func (c *Command) monitor() {
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
