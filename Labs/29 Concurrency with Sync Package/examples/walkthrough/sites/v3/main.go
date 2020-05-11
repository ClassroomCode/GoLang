// section: code
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	args := []string{"http://youtube.com", "http://en.wikipedia.org", "http://facebook.com"}

	now := time.Now()

	for _, a := range args {
		get(a)
	}

	fmt.Printf("It took %s for the entire process to finish\n", time.Since(now))
}

func get(s string) {
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
2019/10/09 16:56:53 Retrieved site http://youtube.com in 662.923755ms
2019/10/09 16:56:54 Retrieved site http://en.wikipedia.org in 397.967866ms
2019/10/09 16:56:54 Retrieved site http://facebook.com in 459.329112ms
It took 1.520455229s for the entire process to finish
// section: output

*/
