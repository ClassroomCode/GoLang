package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const N = 50

// section: wg
var found string
var moot = &sync.Mutex{}

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < N; i++ {
		wg.Add(1)
		go search(i, wg)
	}
	wg.Wait()

	fmt.Print(found)
}

// section: wg

// section: mu
func search(i int, wg *sync.WaitGroup) {
	for {
		fmt.Print(".")
		moot.Lock()
		if found != "" {
			moot.Unlock()
			wg.Done()
			break
		}
		moot.Unlock()
		x := rand.Intn(N)
		if x == 42 {
			moot.Lock()
			fmt.Print("found")
			found = fmt.Sprintf("routine %d: found 42!", i)
			moot.Unlock()
			wg.Done()
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
}

// section: mu
