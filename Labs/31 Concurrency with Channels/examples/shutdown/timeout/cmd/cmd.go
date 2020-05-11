package cmd

import (
	"log"
	"sync"
	"time"
)

func New() *Command {
	return &Command{
		wg:   &sync.WaitGroup{},
		quit: make(chan struct{}),
	}
}

type Command struct {
	wg   *sync.WaitGroup
	quit chan struct{}
}

func (c *Command) Open() error {
	c.wg.Add(1)
	go c.monitor()

	return nil
}

func (c *Command) Close() error {
	close(c.quit)
	c.wg.Wait()
	// add artificial shutdown time for example
	time.Sleep(5 * time.Second)
	return nil
}

func (c *Command) monitor() {
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
