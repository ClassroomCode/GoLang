package main

import "fmt"

func main() {
	a1 := [2]string{"one", "two"}
	a2 := [2]string{}

	a2 = a1

	fmt.Println(a2)
	a3 := [3]string{}

	// This can't be done, as it is not of the same type
	//a3 = a2

	fmt.Println(a3)
}
