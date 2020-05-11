package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if err := fileCopy("readme.txt", "readme-copy.txt"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("file copied successfully")
}

func fileCopy(sourceName string, destName string) error {
	src, err := os.Open(sourceName)
	if err != nil {
		return err
	}
	defer func() {
		fmt.Println("closing", sourceName)
		src.Close()
	}()

	dst, err := os.Create(destName)
	if err != nil {
		return err
	}
	defer func() {
		fmt.Println("closing", destName)
		dst.Close()
	}()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}
	return nil
}
