package main

import "fmt"

func main() {
	a := make([]int, 1, 3)

	fmt.Println(a)      // [0]
	fmt.Println(len(a)) // 1
	fmt.Println(cap(a)) // 3
}
