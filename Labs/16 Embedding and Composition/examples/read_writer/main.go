package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type RW struct {
	data []byte
}

func (rw *RW) Read(p []byte) (int, error) {
	n := copy(p, rw.data)
	return n, nil
}

func (rw *RW) Write(p []byte) (int, error) {
	rw.data = p
	return len(rw.data), nil
}

func readWriter(rw io.ReadWriter) error {
	_, err := rw.Write([]byte("Hello world!"))
	if err != nil {
		return err
	}
	if s, ok := rw.(io.Seeker); ok {
		s.Seek(io.SeekStart, 0)
	}
	b := make([]byte, 512)
	_, err = rw.Read(b)
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}

func main() {
	err := readWriter(&RW{data: []byte{}})
	if err != nil {
		log.Fatal(err)
	}
	t, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatal(err)
	}
	err = readWriter(t)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
}
