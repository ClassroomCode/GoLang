package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// section: nop
func main() {

	r := strings.NewReader("Hello, World!")

	if err := read(ioutil.NopCloser(r)); err != nil {
		fmt.Println(err)
	}
}

func read(r io.ReadCloser) error {
	if _, err := io.Copy(os.Stdout, r); err != nil {
		return err
	}
	return r.Close()
}

// section: nop
