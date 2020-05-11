package main

import "fmt"

func main() {
	x := "George"
	printValue(&x)
	fmt.Println(x)

	s := new(string)
	*s = "hello"
	fmt.Println(s, *s)
}

func printValue(s *string) {
	fmt.Println(s)
	fmt.Println(*s)
	*s = "Ringo"
}
