package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	a := []string{}
	a[42] = "Bring a towel"
}
