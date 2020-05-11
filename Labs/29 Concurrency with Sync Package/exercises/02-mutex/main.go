package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// section: template
type Numbers struct {
	// Add a sync.RWMutex as a field (not embedded)
	data map[int]int
}

func (n *Numbers) Find(key int) (int, error) {
	// Protect this with a read lock
	k, ok := n.data[key]
	if !ok {
		return 0, fmt.Errorf("number not found: %d", key)
	}
	return k, nil
}

func (n *Numbers) Load() {
	var i int
	for {
		// Protect this with a write lock
		n.data[i] = i
		i++
		randomSleep(75) // simulate real world load
	}
}

func NewNumbers() *Numbers {
	n := &Numbers{
		data: map[int]int{},
	}
	go n.Load()
	return n
}

var wg sync.WaitGroup

func main() {
	n := NewNumbers()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			for {
				_, err := n.Find(i)
				if err == nil {
					fmt.Printf("found %d\n", i)
					break
				}
				fmt.Println(err)
				randomSleep(25) // create some delay in our goroutine
			}
		}
	}()
	wg.Wait()
}

// section: template

func randomSleep(i int) {
	// sleep for a random amount of time to add "chaos"
	time.Sleep(time.Duration(rand.Intn(i)) * time.Millisecond)
}
