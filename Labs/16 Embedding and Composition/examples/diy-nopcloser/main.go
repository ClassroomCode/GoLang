package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// section: nop-closer
type Closer struct {
	io.Reader
}

func (c Closer) Close() error {
	return nil
}

// section: nop-closer

// section: nop-implemented
func main() {
	c := Closer{
		Reader: strings.NewReader("Hello, World!"),
	}
	if err := read(c); err != nil {
		fmt.Println(err)
	}

}

func read(r io.ReadCloser) error {
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return r.Close()
}

// section: nop-implemented
