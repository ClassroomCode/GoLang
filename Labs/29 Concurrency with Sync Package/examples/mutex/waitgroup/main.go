package main

import (
	"fmt"
	"sync"
)

// section: wg
type WaitGroup struct {
	mu      sync.Mutex
	quit    chan (struct{})
	counter int
}

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

// section: wg

// section: main
func main() {
	var wg WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("hi")
		}()
	}

	wg.Wait()
}

// section: main
