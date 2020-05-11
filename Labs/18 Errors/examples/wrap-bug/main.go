package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

// section: code
func main() {
	// Check for the sentinel error of `io.EOF`
	err := cat("./README.md")
	if !errors.Is(err, io.EOF) {
		log.Fatalf("something went unexpected: %s", err)
	}
}

func cat(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	buff := make([]byte, 10)

	for {
		i, err := f.Read(buff)
		if err != nil {
			return fmt.Errorf("got a read error: %s", err)
		}
		if _, err := os.Stdout.Write(buff[:i]); err != nil {
			return fmt.Errorf("write error: %s", err)
		}
	}
}

// section: code

/*
// section: output
# Title

Description of project

## Getting Started

Quick start guides here...

### Prerequisites

Here is what you need to know before you start...
2020/04/18 15:13:56 something went unexpected: got a read error: EOF
exit status 1
// section: output
*/
