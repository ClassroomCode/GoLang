package main

import (
	"io"
	"log"
	"os"
)

// section: code
func main() {
	f, err := os.Open("./README.md")
	if err != nil {
		log.Fatal(err)
	}

	buff := make([]byte, 10)

	for {
		i, err := f.Read(buff)
		if err != nil {
			// Check for the sentinel error of `io.EOF`
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
		if _, err := os.Stdout.Write(buff[:i]); err != nil {
			log.Fatal(err)
		}
	}
}

// section: code
