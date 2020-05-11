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

func get(s string, responses chan resp) {
	now := time.Now()

	_, err := http.Get(s)
	if err != nil {
		responses <- resp{site: s, err: err}
		return
	}
	responses <- resp{site: s, duration: time.Since(now)}
}

// section: code
