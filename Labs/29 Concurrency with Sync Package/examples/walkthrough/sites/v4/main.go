// section: code
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	args := []string{"http://youtube.com", "http://en.wikipedia.org", "http://facebook.com"}

	wg := &sync.WaitGroup{}
	wg.Add(len(args))

	now := time.Now()

	for _, a := range args {
		go get(a, wg)
	}

	wg.Wait()

	fmt.Printf("It took %s for the entire process to finish\n", time.Since(now))
}

func get(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	now := time.Now()

	_, err := http.Get(s)
	if err != nil {
		log.Printf("failed to retrieve %s: %s\n", s, err)
		return
	}
	log.Printf("Retrieved site %s in %s\n", s, time.Since(now))
}

// section: code

/*
// section: output
2019/10/09 17:02:40 Retrieved site http://en.wikipedia.org in 420.20124ms
2019/10/09 17:02:40 Retrieved site http://youtube.com in 474.16273ms
2019/10/09 17:02:41 Retrieved site http://facebook.com in 1.014454619s
It took 1.014520896s for the entire process to finish
// section: output

*/
