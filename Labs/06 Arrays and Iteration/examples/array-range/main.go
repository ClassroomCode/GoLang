package main

import "fmt"

func main() {
	names := [4]string{"John", "Paul", "George", "Ringo"}

	for i, n := range names {
		fmt.Printf("%d - %s\n", i, n)
	}
}
