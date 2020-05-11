package main

import "fmt"

func main() {

	// section: do
	var i int
	for {
		fmt.Println(i)
		i += 2
		if i >= 3 {
			break
		}
	}
	// section: do

}
