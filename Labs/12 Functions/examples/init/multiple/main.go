// section: code
package main

import "fmt"

func init() {
	fmt.Println("First init")
}

func init() {
	fmt.Println("Second init")
}

func init() {
	fmt.Println("Third init")
}

func init() {
	fmt.Println("Fourth init")
}

func main() {}

// section: code

/*
// section: output
First init
Second init
Third init
Fourth init
// section: output
*/
