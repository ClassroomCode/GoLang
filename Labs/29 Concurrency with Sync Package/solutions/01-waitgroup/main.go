package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	wg := &sync.WaitGroup{} // <- Add a waitgroup

	root := "./files"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			wg.Add(1)              // <-- Track each new go routine
			go charCount(path, wg) // <-- Launch in a go routine, pass the WaitGroup
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	wg.Wait() // <-- Wait for all go routines to finish
	fmt.Println("process took", time.Since(start))
}

func charCount(path string, wg *sync.WaitGroup) { // <-- Accept the WaitGroup as an argument
	defer wg.Done() // <-- be sure to mark the go routine as done
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("error opening %s: %s\n", err)
		return
	}
	defer f.Close()
	var count int
	buff := make([]byte, bytes.MinRead)
	for {
		i, err := f.Read(buff)
		count += i
		if err == io.EOF {
			fmt.Printf("%s has %d characters\n", path, count)
			return
		}
	}
}
