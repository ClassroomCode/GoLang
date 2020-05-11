package main

import (
	"io"
	"os"
)

func main() {
	if err := write("readme.txt", "This is a readme file"); err != nil {
		panic(err)
	}
}

func write(fileName string, text string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	_, err = io.WriteString(file, "This is a readme file")
	if err != nil {
		file.Close()
		return err
	}
	file.Close()
	return nil
}
