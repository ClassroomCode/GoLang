package main

import (
	"log"
	"net/http"
)

func main() {
	args := []string{"http://youtube.com", "http://en.wikipedia.org", "http://facebook.com"}

	for _, a := range args {
		_, err := http.Get(a)
		if err != nil {
			log.Printf("failed to retrieve %s: %s\n", a, err)
			continue
		}
		log.Printf("Retrieved site %s\n", a)
	}
}
