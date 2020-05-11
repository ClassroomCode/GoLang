package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	c := NewCommand()

	c.Open()

	wait := make(chan struct{})
	<-wait

	c.Close()
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
