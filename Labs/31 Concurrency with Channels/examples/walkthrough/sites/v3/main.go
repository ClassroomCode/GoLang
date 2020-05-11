// section: code
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// create a custom type to receive our responses
type resp struct {
	site     string
	duration time.Duration
	err      error
}

func main() {
	args := []string{"http://youtube.com", "http://en.wikipedia.org", "http://facebook.com"}

	responses := make(chan resp, len(args))

	now := time.Now()

	for _, a := range args {
		go get(a, responses)
	}

	// read back our responses
	for i := 0; i < cap(responses); i++ {
		r := <-responses
		if r.err != nil {
			log.Printf("failed to retrieve %s: %s\n", r.site, r.err)
			continue
		}
		log.Printf("Retrieved site %s in %s\n", r.site, r.duration)
	}

	fmt.Printf("It took %s for the entire process to finish\n", time.Since(now))
}

// section: get
func get(s string, responses chan resp) {
	now := time.Now()

	// pick a magic timeout duration
	td := 2 * time.Second
	// create a timeout.
	timeout := time.NewTimer(td)
	// this ensures we don't leak a go routine
	defer timeout.Stop()

	// create a channel to to receive the response on locally
	// make sure it's buffered so you don't leak a goroutine
	work := make(chan resp, 1)

	// retrieve the site in a goroutine to avoid blocking
	go func() {
		// retrieve the website
		_, err := http.Get(s)
		work <- resp{site: s, duration: time.Since(now), err: err}
	}()

	select {
	case r := <-work:
		// retrieve the value from our local function and pass it through to our send channel
		responses <- r
	case <-timeout.C:
		responses <- resp{site: s, err: fmt.Errorf("timed out after %s", td)}
	}
}

// section: get

// section: code
