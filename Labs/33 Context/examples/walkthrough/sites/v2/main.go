// section: code
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// create a custom type to receive our responses
type resp struct {
	site     string
	duration time.Duration
	err      error
}

func main() {
	// Load sites from a file
	args := os.Args[1:]
	// Args might be a single string with line returns to delimit sites, break them apart
	if len(args) == 1 {
		args = strings.Split(args[0], "\n")
	}

	// Create a channel to send all sites to
	// We'll create it buffered so that we throttle how much many go routines will be processing at any given time
	sites := make(chan string, 30)

	// create an unbuffered channel to receive on
	// because we don't have contention or latency on the end of our
	// process, there is no need to buffer this channel
	responses := make(chan resp)

	now := time.Now()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// launch the same number of go routines that we have a buffer size for
	for i := 0; i < cap(sites); i++ {
		go get(ctx, sites, responses)
	}

	// feed all of our data into the send channel
	// this channel will block when it's full, and unblock
	// when space becomes available
	// We will do this in a goroutine so we don't block execution
	go func() {
		for _, s := range args {
			sites <- s
		}
		// ensure that any open go routines close down
		close(sites)
	}()

	// read back our responses
	for i := 0; i < len(args); i++ {
		r := <-responses
		if r.err != nil {
			log.Printf("failed to retrieve %s: %s\n", r.site, r.err)
			continue
		}
		log.Printf("Retrieved site %s in %s\n", r.site, r.duration)
	}
	fmt.Printf("It took %s for the entire process to finish\n", time.Since(now))
}

func get(ctx context.Context, sites chan string, responses chan resp) {
	// range through the receiving channel.  When this channel is closed
	// all go routines will terminate
	for s := range sites {
		now := time.Now()

		// pick a magic timeout duration
		td := 2 * time.Second
		// create a timeout.
		ctx, cancel := context.WithTimeout(ctx, td)
		defer cancel()

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
		case <-ctx.Done():
			responses <- resp{site: s, err: fmt.Errorf("timed out after %s", td)}
		}
	}
}

// section: code
