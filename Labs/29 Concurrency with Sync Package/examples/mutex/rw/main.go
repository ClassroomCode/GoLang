package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/markbates/randx"
)

// section: funcs
type WaitGroup struct {
	mu      sync.RWMutex
	quit    chan (struct{})
	counter int
}

func (wg *WaitGroup) Remaining() int {
	wg.mu.RLock()
	defer wg.mu.RUnlock()
	return wg.counter
}

// section: funcs

func (wg *WaitGroup) Done() {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.counter += -1
	if wg.counter <= 0 {
		if wg.quit != nil {
			close(wg.quit)
		}
	}
}

func (wg *WaitGroup) Add(n int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.counter += n
}

func (wg *WaitGroup) Wait() {
	wg.mu.Lock()
	if wg.quit == nil {
		wg.quit = make(chan struct{})
	}
	wg.mu.Unlock()
	<-wg.quit
}

func main() {
	var wg WaitGroup

	go func() {
		for {
			fmt.Println("Remaining: ", wg.Remaining())
			time.Sleep(50 * time.Millisecond)
		}
	}()

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			randx.RandomSleep(50)
			fmt.Println("hi")
		}()
	}

	wg.Wait()
}
