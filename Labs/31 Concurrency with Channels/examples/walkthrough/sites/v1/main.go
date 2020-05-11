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
2020/02/04 06:56:32 Retrieved site http://youtube.com in 366.257249ms
2020/02/04 06:56:32 Retrieved site http://en.wikipedia.org in 120.45097ms
2020/02/04 06:56:33 Retrieved site http://facebook.com in 371.748335ms
It took 858.755882ms for the entire process to finish
// section: output

*/
