package main

import (
	"io"
	"os"
)

func main() {
	writeSomething(os.Stdout)
	// Hello
}

func writeSomething(w io.Writer) {
	w.Write([]byte("Hello"))
}
