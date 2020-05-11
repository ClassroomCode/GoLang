package main

import (
	"fmt"

	"github.com/markbates/randx"
)

// section: wg
type WaitGroup struct {
	quit    chan (struct{})
	counter int
}

func (wg *WaitGroup) Done() {
	wg.counter += -1
	if wg.counter <= 0 {
		if wg.quit != nil {
			close(wg.quit)
		}
	}
}

func (wg *WaitGroup) Add(n int) {
	wg.counter += n
}

func (wg *WaitGroup) Wait() {
	if wg.quit == nil {
		wg.quit = make(chan struct{})
	}
	<-wg.quit
}

// section: wg

// section: main
func main() {
	var wg WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			randx.RandomSleep(50)
			fmt.Println("hi")
		}()
	}

	wg.Wait()
}

// section: main
