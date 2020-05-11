package main

import (
	"fmt"
	"io"
)

// section: scribe
type scribe struct {
	data []byte
}

func (s *scribe) Write(p []byte) (int, error) {
	s.data = p
	return len(p), nil
}

// section: scribe

func main() {
	s := &scribe{}
	writeSomething(s)
	fmt.Println(string(s.data)) // Hello
}

func writeSomething(w io.Writer) {
	w.Write([]byte("Hello"))
}
