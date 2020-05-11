package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

// <-- add a struct type to send responses back on the channel
type result struct {
	path  string
	count int
	err   error
}

func main() {
	start := time.Now()

	paths := []string{}

	root := "./files"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	ch := make(chan result, len(paths)) // <-- create a buffered channel to start execution right away.

	// <-- launch all the processes, use an anonymous go routine to not block execution
	go func() {
		for _, p := range paths {
			go charCount(p, ch) // <-- Launch in a go routine, pass the channel
		}
	}()

	// <-- Create a loop to read back the size of the buffered channel
	for i := 0; i < cap(ch); i++ {
		r := <-ch // <-- get the next available result from the channel
		fmt.Printf("%s has %d characters\n", r.path, r.count)
	}

	fmt.Println("process took", time.Since(start))
}

func charCount(path string, ch chan result) { // <-- Accept the channel as an argument
	f, err := os.Open(path)
	if err != nil {
		// <-- send back the error on a channel
		ch <- result{path: path, count: 0, err: fmt.Errorf("error opening %s: %s\n", err)}
		return
	}
	defer f.Close()
	var count int
	buff := make([]byte, bytes.MinRead)
	for {
		i, err := f.Read(buff)
		count += i
		if err == io.EOF {
			// <-- send back the result on a channel
			ch <- result{path: path, count: count}
			return
		}
	}
}
