package main

import "fmt"

func main() {
	l, c := slicer([]string{"John", "Paul", "George", "Ringo"})
	fmt.Println(l)
	fmt.Println(c)
}

func slicer(s []string) (int, int) {
	return len(s), cap(s)
}
