package main

import "fmt"

const (
	tomato, apple int = iota + 1, iota + 2
	orange, chevy
	ford, _
)

func main() {
	fmt.Printf("tomato =  %v, apple = %v\n", tomato, apple)
	fmt.Printf("orange =  %v, chevy = %v\n", orange, chevy)
	fmt.Printf("ford   =  %v\n", ford)
}
