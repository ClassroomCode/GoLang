// section: code
package main

import (
	"log"
	"net/http"
)

func main() {
	args := []string{"http://youtube.com", "http://en.wikipedia.org", "http://facebook.com"}

	for _, a := range args {
		get(a)
	}
}

func get(s string) {
	_, err := http.Get(s)
	if err != nil {
		log.Printf("failed to retrieve %s: %s\n", s, err)
		return
	}
	log.Printf("Retrieved site %s\n", s)
}

// section: code

/*

// section: output
2019/10/09 16:54:12 Retrieved site http://youtube.com
2019/10/09 16:54:12 Retrieved site http://en.wikipedia.org
2019/10/09 16:54:13 Retrieved site http://facebook.com
// section: output

*/
