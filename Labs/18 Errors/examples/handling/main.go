package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("/path/to/some/content.file")
	if err != nil {
		// open /path/to/some/content.file: no such file or directory
		log.Fatal(err)
	}
}
