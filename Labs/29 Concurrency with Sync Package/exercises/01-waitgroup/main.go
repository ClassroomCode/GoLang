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

func main() {
	start := time.Now()

	root := "./files"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			charCount(path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("process took", time.Since(start))
}

func charCount(path string) {
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
