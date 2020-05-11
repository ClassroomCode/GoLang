package main

import (
	"fmt"
	"image"
	_ "image/png"
	"io"
	"log"
)

func main() {
	fmt.Println("vim-go")
}

// section: decode
func decode(reader io.Reader) image.Rectangle {
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return m.Bounds()
}

// section: decode
